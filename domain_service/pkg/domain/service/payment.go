package service

import (
	"context"
	"time"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"

	"go.uber.org/zap"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/repository"
)

type Payment interface {
	Pay(ctx context.Context, payerID, payeeID int64, amount int64) error
}

func NewPayment(
	lg *zap.Logger,
	repo repository.Payment,
	bs Balance,
) Payment {
	return &payment{
		logger:         lg,
		repository:     repo,
		balanceService: bs,
	}
}

type payment struct {
	logger         *zap.Logger
	repository     repository.Payment
	balanceService Balance
}

func (p *payment) Pay(ctx context.Context, payerID, payeeID, amount int64) error {
	if err := p.balanceService.Transfer(ctx, payerID, payeeID, amount); err != nil {
		go p.saveFailedHistory(context.Background(), payerID, payeeID, amount)
		return err
	}
	go p.saveSucceededHistory(ctx, payerID, payeeID, amount)
	return nil
}

func (p *payment) saveFailedHistory(ctx context.Context, payerID, payeeID, amount int64) error {
	err := p.repository.SaveHistory(ctx, &entity.PaymentHistory{
		Amount:    amount,
		PayerID:   payerID,
		PayeeID:   payeeID,
		Status:    entity.PaymentStatusFailed,
		CreatedAt: time.Now(),
	})
	if err != nil {
		p.logger.Error(
			"決済失敗の履歴登録に失敗しました",
			zap.Error(err),
			zap.Int64("payerID", payerID),
			zap.Int64("payeeID", payeeID),
			zap.Int64("amount", amount),
		)
		return err
	}
	return nil
}

func (p *payment) saveSucceededHistory(ctx context.Context, payerID, payeeID, amount int64) error {
	err := p.repository.SaveHistory(ctx, &entity.PaymentHistory{
		Amount:    amount,
		PayerID:   payerID,
		PayeeID:   payeeID,
		Status:    entity.PaymentStatusSucceeded,
		CreatedAt: time.Now(),
	})
	if err != nil {
		p.logger.Error(
			"決済成功の履歴登録に失敗しました",
			zap.Error(err),
			zap.Int64("payerID", payerID),
			zap.Int64("payeeID", payeeID),
			zap.Int64("amount", amount),
		)
		return err
	}
	return nil
}
