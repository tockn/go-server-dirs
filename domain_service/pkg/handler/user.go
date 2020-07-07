package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tockn/go-dirs/domain_service/pkg/view"

	"github.com/gorilla/mux"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userID"]

	ctx := r.Context()
	u, err := h.userRepository.GetByID(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(view.UserFromModel(u)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
