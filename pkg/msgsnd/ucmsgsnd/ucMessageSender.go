package ucmsgsnd

import (
	"context"
	"time"

	"github.com/mavr/anonymous-mail/pkg/anonbot"
	"github.com/mavr/anonymous-mail/pkg/msgsnd"
	"github.com/sirupsen/logrus"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Configuration struct
type Configuration struct {
	// Number of workers
	Jobs int
}

// Usecase - this object receive new messages from telegram api and processing it.
type Usecase struct {
	conf Configuration
	repo msgsnd.Repository
	bot  *anonbot.Bot
}

// New create message receiver object.
func New(repo msgsnd.Repository, bot *anonbot.Bot, c Configuration) *Usecase {
	return &Usecase{
		conf: c,
		repo: repo,
		bot:  bot,
	}
}

// Processing messages from queue.
func (uc *Usecase) Processing(ctx context.Context) error {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil

		case <-t.C:
			m, err := uc.repo.GetMessage()
			if err != nil {
				continue
			}

			chat, err := uc.repo.GetChat(m.To)
			if err != nil {
				logrus.WithError(err).WithField("username", m.To).Debug("chat not found")
				continue
			}

			msg := tgbotapi.NewMessage(chat.ID, m.Text)
			_, err = uc.bot.B.Send(msg)
			if err != nil {
				return err
			}
		}

	}
}
