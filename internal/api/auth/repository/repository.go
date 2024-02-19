package repository

import (
	"context"
	"github/Ndraaa15/fitness-tracker-api/internal/entity"
)

type AuthStoreImpl interface {
	NewStoreClient(useTx bool) (AuthStoreClientImpl, error)
}

type AuthStoreClientImpl interface {
	Rollback() error
	Commit() error
	InsertUser(ctx context.Context, user *entity.User) (*entity.User, error)
}
