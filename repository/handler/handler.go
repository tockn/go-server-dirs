package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tockn/go-dirs/repository/repository"
)

type Handler struct {
	userRepository repository.User
}

func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/{userID}", h.GetUser).Methods(http.MethodGet)
	return r
}
