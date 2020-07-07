package rdb

import (
	"context"
	"database/sql"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
	"github.com/tockn/go-dirs/domain_service/pkg/domain/repository"
)

func NewPaymentRepository(db *sql.DB) repository.Payment {
	return &paymentRepository{
		db: db,
	}
}

type paymentRepository struct {
	db *sql.DB
}

func (p *paymentRepository) SaveHistory(ctx context.Context, h *entity.PaymentHistory) error {
	// TODO いい感じのSQLが入る
	return nil
}
