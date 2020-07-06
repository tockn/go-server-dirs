package rdb

import (
	"context"
	"database/sql"

	"github.com/tockn/go-dirs/service/repository"

	"github.com/tockn/go-dirs/service/model"
)

func NewUserRepository(db *sql.DB) repository.User {
	return &userRepository{db: db}
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	row := r.db.QueryRow(`
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM
			users
		WHERE
			id = ?
	`, id)

	var u model.User
	if err := row.Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) IsFriend(ctx context.Context, id1, id2 string) (bool, error) {
	// TODO
	return false, nil
}
