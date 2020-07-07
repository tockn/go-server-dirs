package service

import (
	"context"
	"fmt"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/repository"
	"go.uber.org/zap"
)

type Balance interface {
	Transfer(ctx context.Context, senderID, receiverID, amount int64) error
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
			if _, err := b.repository.Rollback(ctx); err != nil {
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
	senderBalance.Amount -= amount

	receiverBalance, err := b.repository.GetByUserID(ctx, receiverID)
	if err != nil {
		return err
	}
	receiverBalance.Amount += amount

	err = b.repository.Update(ctx, senderBalance)
	if err != nil {
		lg.Error("出金後残高登録に失敗しました")
		return err
	}
	err = b.repository.Update(ctx, receiverBalance)
	if err != nil {
		lg.Error("入金後残高登録に失敗しました")
		return err
	}

	ctx, err = b.repository.Commit(ctx)
	if err != nil {
		lg.Error("送金処理のCommitに失敗しました")
		return err
	}
	return nil
}
