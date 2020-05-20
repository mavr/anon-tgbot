package ucmsgrecv

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

const (
	cmdStart = "start"
)

func (uc *Usecase) procCommandStart(m *tgbotapi.Message) error {
	return uc.bot.SendStaffMessage(m.From.LanguageCode, staffWellcomeMessage, m.Chat.ID)
}

func (uc *Usecase) procCommandUknown(m *tgbotapi.Message) error {
	return uc.bot.SendStaffMessage(m.From.LanguageCode, staffErrWrongCommand, m.Chat.ID)
}
