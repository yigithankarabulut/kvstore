package basehttphandler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

// Handler defines base handler behaviours.
type Handler struct {
	ServerEnv     string
	Logger        *slog.Logger
	CancelTimeout time.Duration
}

// ContentTypeControl If contentType is not application/json, return error.
func (h *Handler) ContentTypeControl(w http.ResponseWriter) error {
	if w.Header().Get("Content-Type") != "application/json" {
		h.JSON(
			w,
			http.StatusBadRequest,
			map[string]string{"error": "Content-Type header is not application/json"},
		)
		return errors.New("Content-Type header is not application/json")
	}
	return nil
}

// JSON writes json response.
func (h *Handler) JSON(w http.ResponseWriter, status int, d any) {
	j, err := json.Marshal(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, _ = w.Write(j)
}
