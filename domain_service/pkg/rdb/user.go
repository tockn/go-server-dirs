package rdb

import (
	"context"
	"database/sql"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/repository"
)

func NewUserRepository(db *sql.DB) repository.User {
	return &userRepository{db: db}
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) Create(ctx context.Context, name string) (*entity.User, error) {
	res, err := r.db.Exec(`
		INSERT INTO
			users
		VALUES
			(name)
	`, name)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:   id,
		Name: name,
	}, nil
}
