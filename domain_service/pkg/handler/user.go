package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/service"

	"github.com/tockn/go-dirs/domain_service/pkg/view"
)

type User struct {
	Service service.User
}

func (h *User) Create(w http.ResponseWriter, r *http.Request) {
	var req view.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	u, err := h.Service.Create(ctx, req.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := view.CreateUserResponse{
		ID:   u.ID,
		Name: u.Name,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
