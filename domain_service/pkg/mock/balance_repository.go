package mock

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
)

type BalanceRepository struct {
	Began               bool
	Committed           bool
	Rollbacked          bool
	ExpectedError       error
	ExpectedGetByUserID *entity.Balance
}

func (b *BalanceRepository) Begin(ctx context.Context) (context.Context, error) {
	b.Began = true
	return ctx, nil
}

func (b *BalanceRepository) Commit(ctx context.Context) (context.Context, error) {
	b.Committed = true
	return ctx, nil
}

func (b *BalanceRepository) Rollback(ctx context.Context) (context.Context, error) {
	b.Rollbacked = true
	return ctx, nil
}

func (b *BalanceRepository) Update(ctx context.Context, eb *entity.Balance) error {
	return nil
}

func (b *BalanceRepository) GetByUserID(ctx context.Context, userID int64) (*entity.Balance, error) {
	return b.ExpectedGetByUserID, b.ExpectedError
}
