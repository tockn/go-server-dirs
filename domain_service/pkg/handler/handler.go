package handler

import (
	"net/http"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	UserService service.User
}

func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", h.Create).Methods(http.MethodPost)
	return r
}
