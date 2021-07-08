package mock

import (
	"context"
	"github.com/Fantamstick/go-example-api/pkg/domain/entity"
)

type MessageRepo struct {
	messages []entity.Message
}

func NewMockMessageRepository() *MessageRepo {
	var repo = MessageRepo{}

	// generate mock data
	repo.messages = append(repo.messages, entity.Message{
		Id:      1,
		UserId:  1,
		Message: "test message id 1",
	})
	repo.messages = append(repo.messages, entity.Message{
		Id:      2,
		UserId:  1,
		Message: "test message id 2",
	})
	repo.messages = append(repo.messages, entity.Message{
		Id:      3,
		UserId:  2,
		Message: "test message id 3",
	})
	repo.messages = append(repo.messages, entity.Message{
		Id:      4,
		UserId:  2,
		Message: "test message id 4",
	})

	return &repo
}

func (r MessageRepo) ListMessages(ctx context.Context, userId int) ([]entity.Message, error) {
	var messages = make([]entity.Message, 0)

	for _, message := range r.messages {
		if message.UserId == userId {
			messages = append(messages, message)
		}
	}

	return messages, nil
}
