package persistence

import (
	"github.com/Fantamstick/go-example-api/pkg/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	User    repository.IUserRepository
	Message repository.IMessageRepository
}

func NewRepositories(db *mongo.Database) (*Repositories, error) {
	return &Repositories{
		User:    NewUserRepository(db),
		Message: NewMessageRepository(db),
	}, nil
}
