package local

import "github.com/mavr/anonymous-mail/models"

// SaveMessage saving message in local storage
func (s *Store) SaveMessage(m *models.Message) error {
	s.mMessage <- *m

	return nil
}

// GetMessage return message from queue
func (s *Store) GetMessage() (*models.Message, error) {
	select {
	case m := <-s.mMessage:
		return &m, nil

	default:
		return nil, ErrEmptyQueue
	}
}

// GetMessageCh return chan message
func (s *Store) GetMessageCh() chan models.Message {
	return s.mMessage
}
