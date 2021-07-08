package v1

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) GetUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		users, err := h.repo.User.ListUsers(ctx)
		if err != nil {
			msg := "failed to get user"
			http.Error(w, msg, http.StatusInternalServerError)
			h.logger.Error(msg, zap.Error(err))
			return
		}

		b, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.logger.Error("failed to marshal user", zap.Error(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	})
}
