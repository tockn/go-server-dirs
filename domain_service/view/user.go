package view

import (
	"github.com/tockn/go-dirs/domain_service/rdb/model"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func UserFromModel(m *model.User) *User {
	return &User{
		ID:   m.ID,
		Name: m.Name,
	}
}
