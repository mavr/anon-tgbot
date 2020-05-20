package anonbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// TgFielder interface for access to BotAPI fields
type TgFielder interface {
	Self() tgbotapi.User
}
