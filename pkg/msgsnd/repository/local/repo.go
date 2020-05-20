package local

import (
	"github.com/mavr/anonymous-mail/models"
	"github.com/mavr/anonymous-mail/pkg/storage/local"
)

// MsgLocalRepository local storage for messages
type MsgLocalRepository struct {
	m *local.Store
}

// New create local message repository
func New(m *local.Store) *MsgLocalRepository {
	return &MsgLocalRepository{
		m: m,
	}
}

// GetMessage get message from local repository
func (r *MsgLocalRepository) GetMessage() (*models.Message, error) {
	return r.m.GetMessage()
}
