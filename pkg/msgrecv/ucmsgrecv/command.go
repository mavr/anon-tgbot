package ucmsgrecv

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mavr/anonymous-mail/models"
	"github.com/pkg/errors"
)

const (
	cmdStart = "start"
)

func (uc *Usecase) procCommandStart(m *tgbotapi.Message) error {
	if err := uc.chat.NewChatNotificate(&models.Chat{
		ID:      m.Chat.ID,
		UserUID: m.Chat.UserName,
	}); err != nil {
		return errors.Wrap(err, "new chat notification failed")
	}

	return uc.bot.SendStaffMessage(m.From.LanguageCode, staffWellcomeMessage, m.Chat.ID)
}

func (uc *Usecase) procCommandUknown(m *tgbotapi.Message) error {
	return uc.bot.SendStaffMessage(m.From.LanguageCode, staffErrWrongCommand, m.Chat.ID)
}
