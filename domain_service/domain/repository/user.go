package repository

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/domain/entity"
)

type User interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
	IsFriend(ctx context.Context, id1, id2 string) (bool, error)
}
