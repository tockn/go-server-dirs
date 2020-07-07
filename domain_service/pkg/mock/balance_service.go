package mock

import "context"

type BalanceService struct {
	ExpectedError error
}

func (b *BalanceService) Transfer(ctx context.Context, senderID, receiverID, amount int64) error {
	return b.ExpectedError
}
