package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUserByID(db *sql.DB, id string) (*User, error) {
	row := db.QueryRow(`
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

	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}
