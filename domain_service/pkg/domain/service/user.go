package service

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
)

type User interface {
	Create(ctx context.Context, name string) (*entity.User, error)
}
