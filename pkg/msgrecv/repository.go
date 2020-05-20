package msgrecv

import "github.com/mavr/anonymous-mail/models"

// Repository interface
type Repository interface {
	SaveMessage(*models.Message) error
	SetChat(c *models.Chat) error
}
