package msgrecv

import "errors"

var (
	// ErrWrongFormat received message has wrong format
	ErrWrongFormat = errors.New("message has wrong format")
)