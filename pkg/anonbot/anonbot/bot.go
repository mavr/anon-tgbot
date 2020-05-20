package anonbot

import (
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mavr/anonymous-mail/pkg/anonbot"
	"github.com/mavr/anonymous-mail/pkg/config"
)

// Bot wrap fot telegram bot
type Bot struct {
	*tgbotapi.BotAPI
}

// New create new Bot structure and configurate telegram bot, check connection
func New(conf config.Config) (anonbot.AnonBot, error) {
	httpCli := &http.Client{}
	return newAnonBotWithClient(conf, httpCli)
}

func newAnonBotWithClient(conf config.Config, cli *http.Client) (anonbot.AnonBot, error) {
	bot, err := tgbotapi.NewBotAPIWithClient(conf.Bot.TGBotToken, cli)
	if err != nil {
		return nil, err
	}

	// TODO : add configure buffer size
	// bot.Buffer = uc.conf.NumberJobs

	return &Bot{
		BotAPI: bot,
	}, err
}

// Self return self field
func (b *Bot) Self() tgbotapi.User {
	return b.BotAPI.Self
}

// SendStaffMessage send message to chat with chosen language translate.
// Message would be delivery as is in text param.
func (b *Bot) SendStaffMessage(langCode string, text anonbot.Staff, chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, text.Get(langCode))
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := b.Send(msg)

	return err
}

// SendStaffMessageWithTitle send formatted message to chat with chosen language translate.
// Message has format:
// Bold Title
//
// simple message text
func (b *Bot) SendStaffMessageWithTitle(langCode string, title, text anonbot.Staff, chatID int64) error {
	t := fmt.Sprintf("*%s*\n\n%s", title.Get(langCode), text.Get(langCode))
	msg := tgbotapi.NewMessage(chatID, t)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := b.Send(msg)

	return err
}

// SendMessageWithTitle send message with as is text, but formatted title.
func (b *Bot) SendMessageWithTitle(langCode string, title anonbot.Staff, text string, chatID int64) error {
	t := fmt.Sprintf("*%s*\n\n%s", title.Get(langCode), text)
	msg := tgbotapi.NewMessage(chatID, t)
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := b.Send(msg)

	return err
}
