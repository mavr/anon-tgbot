package local

import "errors"

var (
	// ErrNotFound entity not found in local databaseS
	ErrNotFound = errors.New("not found")

	// ErrEmptyQueue queue is empty
	ErrEmptyQueue = errors.New("empty queue")
)
