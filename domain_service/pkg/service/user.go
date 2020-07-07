package service

import (
	"context"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/repository"

	"github.com/tockn/go-dirs/domain_service/pkg/rdb/model"
)

type User interface {
	// GetByIDAndViewerIDは、userIDとviewerIDがfriendの場合userIDのUser情報を返し、
	// friendではない場合Nameを匿名にしたuserIDのUser情報を返します
	GetByIDAndViewerID(ctx context.Context, userID, viewerID string) (*model.User, error)
}

func NewUser(repo repository.User) User {
	return &user{
		repository: repo,
	}
}

type user struct {
	repository repository.User
}

func (s *user) GetByIDAndViewerID(ctx context.Context, userID, viewerID string) (*model.User, error) {
	u, err := s.repository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	isFriend, err := s.repository.IsFriend(ctx, userID, viewerID)
	if err != nil {
		return nil, err
	}

	if !isFriend {
		u.Name = model.AnonymousName
	}
	return u, nil
}
