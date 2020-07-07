package repository

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
)

type User interface {
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
