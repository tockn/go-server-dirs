package mock

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
)

type PaymentRepository struct {
	ExpectedError error
}

func (p *PaymentRepository) SaveHistory(ctx context.Context, h *entity.PaymentHistory) error {
	return p.ExpectedError
}
