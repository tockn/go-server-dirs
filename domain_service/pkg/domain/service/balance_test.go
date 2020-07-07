package service

import (
	"context"
	"testing"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"

	"github.com/stretchr/testify/assert"

	"github.com/tockn/go-dirs/domain_service/pkg/mock"
	"go.uber.org/zap"
)

func TestBalance_Transfer(t *testing.T) {
	type fields struct {
		senderID   int64
		receiverID int64
		amount     int64
	}
	tests := []struct {
		name        string
		fields      fields
		repo        *mock.BalanceRepository
		began       bool
		committed   bool
		rollbacked  bool
		isExpectErr bool
	}{
		{
			name: "success",
			fields: fields{
				senderID:   0,
				receiverID: 1,
				amount:     100,
			},
			repo:       &mock.BalanceRepository{ExpectedGetByUserID: balance1()},
			began:      true,
			committed:  true,
			rollbacked: false,
		},
		{
			name: "残高不足の時",
			fields: fields{
				senderID:   0,
				receiverID: 1,
				amount:     balance1().Amount + 100,
			},
			repo:        &mock.BalanceRepository{ExpectedGetByUserID: balance1()},
			began:       true,
			committed:   false,
			rollbacked:  true,
			isExpectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewBalance(zap.NewNop(), tt.repo)
			if err := s.Transfer(context.Background(), tt.fields.senderID, tt.fields.receiverID, tt.fields.amount); err != nil {
				if !tt.isExpectErr {
					t.Fatal(err)
				}
				t.Skip()
			}
			assert.Equal(t, tt.began, tt.repo.Began)
			assert.Equal(t, tt.committed, tt.repo.Committed)
			assert.Equal(t, tt.rollbacked, tt.repo.Rollbacked)
		})
	}
}

func balance1() *entity.Balance {
	return &entity.Balance{
		Amount: 100,
		UserID: 0,
	}
}
