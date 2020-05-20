package anonbot

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mavr/anonymous-mail/pkg/config"
)

// Bot wrap fot telegram bot
type Bot struct {
	B *tgbotapi.BotAPI
}

// New create new Bot structure and configurate telegram bot, check connection
func New(conf config.Config) (*Bot, error) {
	httpCli := &http.Client{}
	return newAnonBotWithClient(conf, httpCli)
}

func newAnonBotWithClient(conf config.Config, cli *http.Client) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPIWithClient(conf.Bot.TGBotToken, cli)
	if err != nil {
		return nil, err
	}

	return &Bot{
		B: bot,
	}, err
}
