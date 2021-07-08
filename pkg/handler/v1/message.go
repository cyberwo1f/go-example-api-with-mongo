package v1

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetMessages() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check get path parameter
		param := strings.TrimPrefix(r.URL.Path, "/message/list/")
		userId, err := strconv.Atoi(param)
		if err != nil {
			msg := "failed to get path parameter"
			http.Error(w, msg, http.StatusBadRequest)
			h.logger.Error(msg, zap.Error(err))
			return
		}

		// get messages
		ctx := context.Background()
		messages, err := h.repo.Message.ListMessages(ctx, userId)
		if err != nil {
			msg := "failed to get messages"
			http.Error(w, msg, http.StatusInternalServerError)
			h.logger.Error(msg, zap.Error(err))
			return
		}

		// return json
		b, err := json.Marshal(messages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.logger.Error("failed to marshal messages", zap.Error(err))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
		return
	})
}
