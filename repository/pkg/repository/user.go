package repository

import (
	"context"

	"github.com/tockn/go-dirs/repository/pkg/model"
)

type User interface {
	Create(ctx context.Context, name string) (*model.User, error)
}
