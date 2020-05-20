package local

import "github.com/mavr/anonymous-mail/models"

const (
	messageBuffer = 64
)

// Store local store
type Store struct {
	// mMessage map[string]models.Message
	mMessage chan (models.Message)
	mChat    map[string]models.Chat
}

// New create new local store
func New() (*Store, error) {
	return &Store{
		mMessage: make(chan models.Message, messageBuffer),
		mChat:    make(map[string]models.Chat),
	}, nil
}
