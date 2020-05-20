package msgsnd

import "github.com/mavr/anonymous-mail/models"

// Repository iterface
type Repository interface {
	GetMessage() (*models.Message, error)
	GetChat(username string) (*models.Chat, error)
}
