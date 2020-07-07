package mock

import (
	"context"
)

type PaymentService struct {
	ExpectedError error
}

func (p *PaymentService) Pay(ctx context.Context, payerID, payeeID int64, amount int64) error {
	return p.ExpectedError
}
