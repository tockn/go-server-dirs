package mock

import (
	"context"

	"github.com/tockn/go-dirs/repository/pkg/model"
)

type UserRepository struct {
	ExpectedUser  *model.User
	ExpectedError error
}

func (r *UserRepository) Create(ctx context.Context, id string) (*model.User, error) {
	return r.ExpectedUser, r.ExpectedError
}
