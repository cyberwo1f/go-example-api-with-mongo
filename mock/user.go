package mock

import (
	"context"
	"github.com/Fantamstick/go-example-api/pkg/domain/entity"
	"github.com/Fantamstick/go-example-api/pkg/domain/repository"
)

type UserRepo struct {
	users []entity.User
}

var _ repository.IUserRepository = (*UserRepo)(nil)

func NewMockUserRepository() *UserRepo {
	var repo = UserRepo{}

	// generate mock data
	repo.users = append(repo.users, entity.User{
		Id:   1,
		Name: "Hoge",
	})
	repo.users = append(repo.users, entity.User{
		Id:   2,
		Name: "Fuga",
	})

	return &repo
}

func (r UserRepo) ListUsers(ctx context.Context) ([]entity.User, error) {
	return r.users, nil
}
