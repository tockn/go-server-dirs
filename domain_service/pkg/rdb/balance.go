package rdb

import (
	"context"
	"database/sql"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
	"github.com/tockn/go-dirs/domain_service/pkg/domain/repository"
)

func NewBalanceRepository(db *sql.DB) repository.Balance {
	return &balanceRepository{
		db: db,
	}
}

type balanceRepository struct {
	db *sql.DB
}

func (b balanceRepository) Begin(ctx context.Context) (context.Context, error) {
	// TODO いい感じのSQLが入る
	return ctx, nil
}

func (b balanceRepository) Commit(ctx context.Context) (context.Context, error) {
	// TODO いい感じのSQLが入る
	return ctx, nil
}

func (b balanceRepository) Rollback(ctx context.Context) (context.Context, error) {
	// TODO いい感じのSQLが入る
	return ctx, nil
}

func (b balanceRepository) GetByUserID(ctx context.Context, userID int64) (*entity.Balance, error) {
	// TODO いい感じのSQLが入る
	return nil, nil
}

func (b balanceRepository) Update(ctx context.Context, eb *entity.Balance) error {
	// TODO いい感じのSQLが入る
	return nil
}
