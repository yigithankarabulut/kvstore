package kvstorehandler

import (
	"context"
	"errors" //nolint:gofumpt
	"net/http"

	"github.com/yigithankarabulut/kvstore/src/internal/kverror" //nolint:gci
)

func (h *kvstoreHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.JSON(
			w,
			http.StatusMethodNotAllowed,
			map[string]string{"error": "method " + r.Method + " not allowed"},
		)
		return
	}
	// verify query params
	if len(r.URL.Query()) == 0 {
		h.JSON(
			w,
			http.StatusNotFound,
			map[string]string{"error": "key query param required"},
		)
		return
	}

	// verify valid key?
	// /?key=foo
	keys, ok := r.URL.Query()["key"]
	if !ok {
		h.JSON(
			w,
			http.StatusNotFound,
			map[string]string{"error": "key not present"},
		)
		return
	} else if len(keys) > 1 {
		h.JSON(
			w,
			http.StatusNotFound,
			map[string]string{"error": "only one key query parameter is required"}, //
		)
		return
	}

	key := keys[0]

	ctx, cancel := context.WithTimeout(r.Context(), h.CancelTimeout)
	defer cancel()

	serviceResponse, err := h.service.Get(ctx, key)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			h.JSON(
				w,
				http.StatusGatewayTimeout,
				map[string]string{"error": err.Error()},
			)
		}

		var kvErr *kverror.Error
		if errors.As(err, &kvErr) {
			clientMessage := kvErr.Message
			if kvErr.Data != nil {
				data, ok := kvErr.Data.(string)
				if ok {
					clientMessage += ", " + data
				}
			}
			if kvErr.Loggable {
				h.Logger.Error("kvstorehandler Get service.Get", "err", clientMessage)
			}

			if kvErr == kverror.ErrKeyNotFound {
				h.JSON(
					w,
					http.StatusNotFound,
					map[string]string{"error": clientMessage})
				return
			}
		}
		h.JSON(
			w,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}
	handlerResponse := &ItemResponse{
		Key:   serviceResponse.Key,
		Value: serviceResponse.Value,
	}
	h.JSON(
		w,
		http.StatusOK,
		handlerResponse,
	)
}
