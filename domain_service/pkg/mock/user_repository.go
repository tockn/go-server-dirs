package mock

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"
)

type UserRepository struct {
	ExpectedUser      *entity.User
	ExpectedUserError error

	ExpectedFriend      bool
	ExpectedFriendError error
}

func (r *UserRepository) GetByID(ctx context.Context, id string) (*entity.User, error) {
	return r.ExpectedUser, r.ExpectedUserError
}

func (r *UserRepository) IsFriend(ctx context.Context, id1, id2 string) (bool, error) {
	return r.ExpectedFriend, r.ExpectedFriendError
}
