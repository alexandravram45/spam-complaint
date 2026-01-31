package api

import (
	"encoding/json"
	"net/http"

	"github.com/alexandravram45/spam-complaint/internal/complaint"
	"github.com/alexandravram45/spam-complaint/internal/detector"
)

type Handler struct {
	service  *complaint.Service
	detector *detector.SpamDetector
}

func NewHandler(s *complaint.Service, d *detector.SpamDetector) *Handler {
	return &Handler{service: s, detector: d}
}

func (h *Handler) Complaints(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req struct {
			Message string `json:"message"`
		}
		_ = json.NewDecoder(r.Body).Decode(&req)

		isSpam := h.detector.IsSpam(req.Message)
		h.service.HandleMessage(req.Message, isSpam)

		w.WriteHeader(http.StatusCreated)

	case http.MethodGet:
		_ = json.NewEncoder(w).Encode(h.service.All())

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
