package msgrecv

import "context"

// Usecase interface
type Usecase interface {
	Processing(ctx context.Context) error
}
