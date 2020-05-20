package ucmsgrecv

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/mavr/anonymous-mail/models"
	"github.com/mavr/anonymous-mail/pkg/anonbot"
	"github.com/mavr/anonymous-mail/pkg/msgrecv"
)

const (
	// Default buffer size for request to telegam api
	defaultUpdatesBufferSize = 4
	defaultUpdateTimer       = 3 * time.Second
)

// Configuration struct
type Configuration struct {
	// Number of workers
	NumberJobs int

	// Buffer size for updates in each request to telegram api
	UpdatesBufferSize int

	Debug bool
}

// Usecase - this object receive new messages from telegram api and processing it.
type Usecase struct {
	conf Configuration
	repo msgrecv.Repository
	bot  *anonbot.Bot
}

// New create message receiver object.
func New(repo msgrecv.Repository, bot *anonbot.Bot, c Configuration) *Usecase {
	if c.UpdatesBufferSize == 0 {
		c.UpdatesBufferSize = c.NumberJobs * defaultUpdatesBufferSize
	}

	return &Usecase{
		conf: c,
		repo: repo,
		bot:  bot,
	}
}

// Processing messages from teleram api.
func (uc *Usecase) Processing(ctx context.Context) error {
	up := tgbotapi.NewUpdate(0)
	up.Timeout = 1

	uc.bot.B.Buffer = uc.conf.NumberJobs

	wg := &sync.WaitGroup{}
	ch := make(chan tgbotapi.Update, uc.conf.UpdatesBufferSize)
	for i := 0; i < uc.conf.NumberJobs; i++ {
		go func(wg *sync.WaitGroup, ch <-chan tgbotapi.Update) {
			for u := range ch {
				func() {
					wg.Add(1)
					defer wg.Done()

					if err := uc.procUpdate(u); err != nil {
						if uc.conf.Debug {
							b, _ := json.Marshal(u)
							fmt.Printf("failed message %s\n", b)
						}
						logrus.WithError(err).
							WithField("update_id", u.UpdateID).
							Debug("processing update failed")

						err = errors.Unwrap(err)

						if errors.Is(err, msgrecv.ErrWrongFormat) {
							// send error to chat
							if u.Message != nil && u.Message.Chat != nil {
								uc.sendStaffMessageWithTitle(
									u.Message.From.LanguageCode,
									staffErrNotDeliver,
									staffWrongMessageFormat,
									u.Message.Chat.ID,
								)

								return
							}

							logrus.WithField("update_id", u.UpdateID).Debug("notify user failed")

							return
						}
					}
				}()
			}
		}(wg, ch)
	}

	t := time.NewTimer(defaultUpdateTimer)
	for {
		select {
		case <-ctx.Done():
			// uc.bot.B.StopReceivingUpdates()
			t.Stop()
			close(ch)
			wg.Wait()

			return nil

		case <-t.C:
			updates, err := uc.bot.B.GetUpdates(up)
			if err != nil {
				return err
			}
			for _, u := range updates {
				ch <- u
				up.Offset = u.UpdateID + 1
			}

			t.Reset(defaultUpdateTimer)
		}
	}
}

func (uc *Usecase) procUpdate(u tgbotapi.Update) error {
	if u.Message == nil {
		return errors.New("nil message")
	}

	if u.Message.IsCommand() {
		logrus.WithField("chat_id", u.Message.Chat.ID).WithField("command", u.Message.Text).Debug("receive command")

		return uc.procCommand(u.Message)
	}

	logrus.WithField("chat_id", u.Message.Chat.ID).WithField("text", u.Message.Text).Debug("receive message")

	return uc.procMessage(u.Message)
}

func (uc *Usecase) procCommand(m *tgbotapi.Message) error {
	cmd := m.Text[1:]
	switch cmd {
	case cmdStart:
		return uc.procCommandStart(m)

	default:
		logrus.WithField("chat_id", m.Chat.ID).Debug("err: receive uknown command")

		return uc.procCommandUknown(m)
	}
}

func (uc *Usecase) procMessage(m *tgbotapi.Message) error {
	to, msg, err := parser(m.Text)
	if err != nil {
		return errors.Wrap(err, "parsing message failed")
	}

	if err := uc.repo.SetChat(&models.Chat{
		ID:      m.Chat.ID,
		UserUID: m.Chat.UserName,
	}); err != nil {
		return errors.Wrap(err, "set chat failed")
	}

	if err = uc.repo.SaveMessage(&models.Message{
		Text:      msg,
		To:        to,
		Processed: false,
		CreatedAt: time.Now(),
	}); err != nil {
		return errors.Wrap(err, "save message failed")
	}

	return nil
}

func (uc *Usecase) sendStaffMessage(langCode string, text staff, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, text.Get(langCode))
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := uc.bot.B.Send(msg)

	return err
}

func (uc *Usecase) sendStaffMessageWithTitle(langCode string, errorMessage staff, reason staff, chatID int64) error {
	text := fmt.Sprintf("*%s*\n\n%s", errorMessage.Get(langCode), reason.Get(langCode))
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := uc.bot.B.Send(msg)

	return err
}

func parser(text string) (to, msg string, err error) {
	parts := strings.SplitN(strings.TrimSpace(text), " ", 2)
	if len(parts) != 2 {
		err = msgrecv.ErrWrongFormat
		return
	}

	if parts[0][0] != '@' {
		err = msgrecv.ErrWrongFormat
		return
	}

	to = parts[0][1:]
	msg = parts[1]

	return
}
