package persistence

import (
	"context"
	"github.com/cyberwo1f/go-example-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"testing"
	"time"
)

var (
	userRepo    *UserRepo
	messageRepo *MessageRepo
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	opts := &options.ClientOptions{}
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(cfg.DB.URL), opts)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	mongoCtx, mongoCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoCancel()

	if err := mongoClient.Connect(mongoCtx); err != nil {
		log.Printf("failed to connect to mongo db: %s\n", err)
		os.Exit(1)
	}

	if err := mongoClient.Ping(mongoCtx, readpref.Primary()); err != nil {
		log.Printf("failed to ping mongo db: %s\n", err)
		os.Exit(1)
	}

	mongoDB := mongoClient.Database(cfg.DB.Database)
	userRepo = NewUserRepository(mongoDB)
	messageRepo = NewMessageRepository(mongoDB)

	res := m.Run()
	os.Exit(res)
}
