package repository

import (
	"context"

	"github.com/tockn/go-dirs/repository/model"
)

type User interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
}
