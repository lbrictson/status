package pkg

import "context"

type StorageBackend interface {
	SaveOperator(ctx context.Context, operator *Operator) error
	ListOperators(ctx context.Context) ([]Operator, error)
	GetOperatorByEmail(ctx context.Context, email string) (Operator, error)
}

type Operator struct {
	Email          string
	HashedPassword string
	Role           string
}
