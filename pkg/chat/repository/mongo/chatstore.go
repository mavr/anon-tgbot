package local

import (
	"github.com/mavr/anonymous-mail/models"
	"github.com/mavr/anonymous-mail/pkg/chat"
	"github.com/mavr/anonymous-mail/pkg/storage/local"
)

// ChatMongoRepository local storage for messages
type ChatMongoRepository struct {
	m *local.Store
}

// New create local message repository
func New(s *local.Store) chat.Repository {
	return &ChatMongoRepository{
		m: s,
	}
}

// SetChat save chat to local repository
func (r *ChatMongoRepository) SetChat(c *models.Chat) error {
	return r.m.SetChat(c)
}

// GetChat get chat by user uid
func (r *ChatMongoRepository) GetChat(userUID string) (*models.Chat, error) {
	return r.m.GetChat(userUID)
}
