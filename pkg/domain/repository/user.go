package repository

import (
	"context"
	"github.com/Fantamstick/go-example-api/pkg/domain/entity"
)

type IUserRepository interface {
	ListUsers(ctx context.Context) ([]entity.User, error)
}
