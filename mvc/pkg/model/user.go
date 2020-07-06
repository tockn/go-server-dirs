package model

import (
	"database/sql"
)

type User struct {
	ID   int64
	Name string
}

func CreateUser(db *sql.DB, name string) (*User, error) {
	res, err := db.Exec(`
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

	return &User{
		ID:   id,
		Name: name,
	}, nil
}
