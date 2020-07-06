package repository

import (
	"context"

	"github.com/tockn/go-dirs/service/model"
)

type User interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	IsFriend(ctx context.Context, id1, id2 string) (bool, error)
}
