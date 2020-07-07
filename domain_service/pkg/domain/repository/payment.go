package repository

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
)

type Payment interface {
	SaveHistory(ctx context.Context, h *entity.PaymentHistory) error
}
