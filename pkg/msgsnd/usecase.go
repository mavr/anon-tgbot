package msgsnd

import "context"

// Usecase interaface
type Usecase interface {
	Processing(ctx context.Context) error
}
