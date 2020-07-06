package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tockn/go-dirs/repository/pkg/repository"
)

type Handler struct {
	UserRepository repository.User
}

func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/{userID}", h.Create).Methods(http.MethodGet)
	return r
}
