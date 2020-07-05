package view

import (
	"time"

	"github.com/tockn/go-dirs/mvc/model"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func UserFromModel(m *model.User) *User {
	return &User{
		ID:        m.ID,
		Name:      m.Name,
		UpdatedAt: m.UpdatedAt,
	}
}
