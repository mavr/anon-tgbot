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

// SaveMessage save message to local repository
func (r *MsgLocalRepository) SaveMessage(m *models.Message) error {
	return r.m.SaveMessage(m)
}
