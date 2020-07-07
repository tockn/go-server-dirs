package repository

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
)

type Balance interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) (context.Context, error)

	GetByUserID(ctx context.Context, userID int64) (*entity.Balance, error)
	Update(ctx context.Context, eb *entity.Balance) error
}
