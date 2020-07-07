package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/service"

	"github.com/tockn/go-dirs/domain_service/pkg/view"
)

type Payment struct {
	Service service.Payment
}

func (h *Payment) Pay(w http.ResponseWriter, r *http.Request) {
	var req view.PayRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.Service.Pay(ctx, req.PayerID, req.PayeeID, req.Amount); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(view.PayResponse{
			Status:        "failed",
			FailedMessage: err.Error(),
		}); err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(view.PayResponse{
		Status: "succeeded",
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
