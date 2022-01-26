package repository

import (
	"context"
	"github.com/cyberwo1f/go-example-api/pkg/domain/entity"
)

type IUserRepository interface {
	ListUsers(ctx context.Context) ([]entity.User, error)
}
