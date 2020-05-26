package chat

import "github.com/mavr/anonymous-mail/models"

// Usecase interface
type Usecase interface {
	// RegisterPublisher create connect for new chat notification.
	RegisterPublisher() error

	// RegisterSubscriber create connect for consumer
	// and start receive messages
	RegisterSubscriber() error

	// NewChatNotificate method for notification about new chat.
	// In this implementation inforamation about new user's chat
	// delivery to service which stored it.
	// Note! For use this notification method RegisterPublisher()
	// must be called first.
	NewChatNotificate(c *models.Chat) error

	// GetNewChatNotificate this is blocking operation which return
	// notification about new chat.
	// Note! There is RegisterSubscriber() need to call before
	// use this method.
	GetNewChatNotificate() (*models.Chat, error)

	// GetNewChatNotificateChan this is blocking operation which return
	// notification about new chat.
	// Note! There is RegisterSubscriber() need to call before
	// use this method.
	GetNewChatNotificateChan() (<-chan *models.Chat, error)

	// SaveNewChat storing new chat.
	SaveNewChat(c *models.Chat) error

	// GetChatByUserUID find chat by user uid.
	GetChatByUserUID(userUID string) (*models.Chat, error)
}
