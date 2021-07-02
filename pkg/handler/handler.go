package handler

import (
	"github.com/Fantamstick/go-example-api/pkg/handler/version"
	"go.uber.org/zap"
)

type Handler struct {
	Version *version.Handler
}

func NewHandler(logger *zap.Logger, ver string) *Handler {
	h := &Handler{
		Version: version.NewHandler(logger.Named("version"), ver),
	}

	return h
}
