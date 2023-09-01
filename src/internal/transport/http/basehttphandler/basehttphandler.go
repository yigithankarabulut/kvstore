package basehttphandler

import (
	"encoding/json"
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
