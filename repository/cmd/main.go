package main

import (
	"database/sql"
	"net/http"

	"github.com/tockn/go-dirs/repository/pkg/handler"
	"github.com/tockn/go-dirs/repository/pkg/rdb"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	db, err := sql.Open("mysql", "")
	if err != nil {
		return err
	}
	userRepo := rdb.NewUserRepository(db)
	h := handler.Handler{UserRepository: userRepo}

	r := h.Router()
	return http.ListenAndServe(":8080", r)
}
