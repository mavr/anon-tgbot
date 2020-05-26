package ucmsgsnd

import (
	"context"

	"github.com/mavr/anonymous-mail/pkg/anonbot"
	"github.com/mavr/anonymous-mail/pkg/chat"
	"github.com/mavr/anonymous-mail/pkg/msgsnd"
	"github.com/sirupsen/logrus"
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
	chat chat.Usecase
	bot  anonbot.AnonBot
}

// New create message receiver object.
func New(repo msgsnd.Repository, chat chat.Usecase, bot anonbot.AnonBot, c Configuration) *Usecase {
	return &Usecase{
		conf: c,
		chat: chat,
		repo: repo,
		bot:  bot,
	}
}

// Processing messages from queue.
func (uc *Usecase) Processing(ctx context.Context) error {
	if err := uc.chat.RegisterSubscriber(); err != nil {
		return err
	}

	chNotificate, _ := uc.chat.GetNewChatNotificateChan()

	chMessage, _ := uc.repo.GetMessageCh()

	for {
		select {
		case <-ctx.Done():
			return nil

		case c := <-chNotificate:
			if err := uc.chat.SaveNewChat(c); err != nil {
				logrus.WithError(err).WithField("chat", c).Error("save chat failed")
			}

			// find stored messages for new client

			// ...

		case m := <-chMessage:
			logrus.WithField("message", m).Info("New message")

			chat, err := uc.repo.GetChat(m.To)
			if err != nil {
				logrus.WithError(err).WithField("username", m.To).Debug("chat not found")
				continue
			}

			if err := uc.bot.SendMessageWithTitle(chat.LangCode, staffNewMessageTitle, m.Text, chat.ID); err != nil {
				logrus.WithError(err).Debug("message send failed")
			}
		}
	}
}
