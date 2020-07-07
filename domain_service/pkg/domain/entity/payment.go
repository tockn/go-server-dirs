package entity

import "time"

type PaymentHistory struct {
	ID        int64
	Amount    int64
	PayerID   int64
	PayeeID   int64
	Status    PaymentStatus
	CreatedAt time.Time
}

type PaymentStatus string

var (
	PaymentStatusSucceeded PaymentStatus = "succeeded"
	PaymentStatusFailed    PaymentStatus = "failed"
)
