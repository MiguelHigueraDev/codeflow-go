package handlers

import (
	"net/http"

	"github.com/miguelhigueradev/codeflow/services/auth-service/internal/httputil"
)

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	httputil.WriteJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
