package mock

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
)

type UserService struct {
	ExpectedUser  *entity.User
	ExpectedError error
}

func (u UserService) Create(ctx context.Context, name string) (*entity.User, error) {
	return u.ExpectedUser, u.ExpectedError
}
