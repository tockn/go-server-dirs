package mock

import (
	"context"

	"github.com/tockn/go-dirs/repository/model"
)

type UserRepository struct {
	ExpectedUser  *model.User
	ExpectedError error
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	return r.ExpectedUser, r.ExpectedError
}
