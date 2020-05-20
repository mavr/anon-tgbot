package local

import (
	"github.com/mavr/anonymous-mail/models"
)

// SetChat complete table of chatid - username
func (s *Store) SetChat(c *models.Chat) error {
	s.mChat[c.UserUID] = *c

	return nil
}

// GetChat return copy of chat with chatuid by username
func (s *Store) GetChat(username string) (*models.Chat, error) {
	c, ok := s.mChat[username]
	if !ok {
		return nil, ErrNotFound
	}

	r := c
	return &r, nil
}
