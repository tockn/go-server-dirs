package service

import (
	"context"
	"fmt"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/repository"
	"go.uber.org/zap"
)

type Balance interface {
	Transfer(ctx context.Context, senderID, receiverID int64, amount int64) error
}

func NewBalance(
	lg *zap.Logger,
	repo repository.Balance,
) Balance {
	return &balance{
		logger:     lg,
		repository: repo,
	}
}

type balance struct {
	logger     *zap.Logger
	repository repository.Balance
}

func (b *balance) Transfer(ctx context.Context, senderID, receiverID int64, amount int64) error {
	lg := b.logger.With(
		zap.Int64("senderID", senderID),
		zap.Int64("receiverID", receiverID),
		zap.Int64("amount", amount),
	)

	ctx, err := b.repository.Begin(ctx)
	if err != nil {
		lg.Error("送金処理のトランザクション開始に失敗しました")
		return err
	}
	defer func() {
		if err != nil {
			if err := b.repository.Rollback(ctx); err != nil {
				lg.Error("送金処理のRollbackに失敗しました")
			}
			return
		}
	}()

	senderBalance, err := b.repository.GetByUserID(ctx, senderID)
	if err != nil {
		return err
	}

	if amount > senderBalance.Amount {
		return fmt.Errorf("残高が足りません")
	}

	if _, err := b.repository.Decrease(ctx, senderID, amount); err != nil {
		lg.Error("出金に失敗しました")
		return err
	}
	if _, err := b.repository.Increase(ctx, receiverID, amount); err != nil {
		lg.Error("入金に失敗しました")
		return err
	}

	if err := b.repository.Commit(ctx); err != nil {
		lg.Error("送金処理のCommitに失敗しました")
		return err
	}
	return nil
}
