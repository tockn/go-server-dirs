package service

import (
	"context"
	"errors"
	"testing"

	"github.com/tockn/go-dirs/domain_service/pkg/mock"
	"go.uber.org/zap"
)

func TestPayment_Pay(t *testing.T) {
	type fields struct {
		senderID   int64
		receiverID int64
		amount     int64
	}
	tests := []struct {
		name           string
		fields         fields
		repo           *mock.PaymentRepository
		balanceService *mock.BalanceService
		isExpectErr    bool
	}{
		{
			name: "success",
			fields: fields{
				senderID:   0,
				receiverID: 1,
				amount:     100,
			},
			repo:           &mock.PaymentRepository{ExpectedError: nil},
			balanceService: &mock.BalanceService{ExpectedError: nil},
		},
		{
			name: "履歴登録でエラーが発生した時",
			fields: fields{
				senderID:   0,
				receiverID: 1,
				amount:     100,
			},
			repo:           &mock.PaymentRepository{ExpectedError: errors.New("save history error")},
			balanceService: &mock.BalanceService{ExpectedError: nil},
			isExpectErr:    true,
		},
		{
			name: "送金処理でエラーが発生した時",
			fields: fields{
				senderID:   0,
				receiverID: 1,
				amount:     100,
			},
			repo:           &mock.PaymentRepository{ExpectedError: nil},
			balanceService: &mock.BalanceService{ExpectedError: errors.New("transfer error")},
			isExpectErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewPayment(zap.NewNop(), tt.repo, tt.balanceService)
			if err := s.Pay(
				context.Background(),
				tt.fields.senderID,
				tt.fields.receiverID,
				tt.fields.amount,
			); err != nil {
				if !tt.isExpectErr {
					t.Fatal(err)
				}
				t.Skip()
			}
		})
	}
}
