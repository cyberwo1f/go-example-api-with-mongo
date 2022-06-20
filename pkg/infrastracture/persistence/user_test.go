package persistence

import (
	"context"
	"github.com/cyberwo1f/go-example-api/pkg/domain/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestUserRepo_ListUsers(t *testing.T) {
	ctx := context.Background()
	// --- 準備 ---
	// 既存を全部消す
	_, _ = userRepo.col.DeleteMany(ctx, bson.D{})
	seeds := []interface{}{
		entity.User{Id: 1, Name: "Hoge"},
		entity.User{Id: 2, Name: "Fuga"},
	}
	_, er := userRepo.col.InsertMany(ctx, seeds)
	assert.NoError(t, er)
	t.Cleanup(func() {
		_, _ = userRepo.col.DeleteMany(ctx, bson.D{})
	})
	// --- 準備 ---

}
