package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	User    User
	Payment Payment
}

func (h *Handler) Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", h.User.Create).Methods(http.MethodPost)

	r.HandleFunc("/payments", h.Payment.Pay).Methods(http.MethodPost)
	return r
}
