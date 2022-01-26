package repository

import (
	"context"
	"github.com/cyberwo1f/go-example-api/pkg/domain/entity"
)

type IMessageRepository interface {
	ListMessages(ctx context.Context, userId int) ([]entity.Message, error)
}
