package v1

import (
	"github.com/cyberwo1f/go-example-api/pkg/infrastracture/persistence"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	repo   *persistence.Repositories
}

func NewHandler(logger *zap.Logger, repositories *persistence.Repositories) *Handler {
	return &Handler{
		logger: logger,
		repo:   repositories,
	}
}
