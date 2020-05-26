package chat

import "github.com/mavr/anonymous-mail/models"

// Repository interface
type Repository interface {
	SetChat(c *models.Chat) error
	GetChat(userUID string) (*models.Chat, error)
}
