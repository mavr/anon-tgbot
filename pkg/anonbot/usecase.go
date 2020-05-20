package anonbot

// AnonBot present methods github.com/go-telegram-bot-api/telegram-bot-api in first case.
// Besides telegram api extends by common methods for send staff messages.
type AnonBot interface {
	// UCTelegramBotAPI telegram-bot-api
	UCTelegramBotAPI

	// Extend methods
	BotWrapper

	// Some TGAPI fields
	TgFielder
}

// BotWrapper extend standart telegram bot api module.
type BotWrapper interface {
	// Group functions for send message to chat in different formates.
	SendStaffMessage(langCode string, text Staff, chatID int64) error
	SendStaffMessageWithTitle(langCode string, title, text Staff, chatID int64) error
	SendMessageWithTitle(langCode string, title Staff, text string, chatID int64) error
}
