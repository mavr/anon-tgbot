package mongo

import "github.com/mavr/anonymous-mail/models"

// Store mongo store
type Store struct {
	mMessage chan (models.Message)
	mChat    map[string]models.Chat
}
