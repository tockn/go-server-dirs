package mock

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
)

type UserRepository struct {
	ExpectedUser  *entity.User
	ExpectedError error
}

func (r *UserRepository) Create(ctx context.Context, name string) (*entity.User, error) {
	return r.ExpectedUser, r.ExpectedError
}
