package server

import (
	"context"
	"encoding/json"
	"github.com/Fantamstick/go-example-api/pkg/config"
	"github.com/Fantamstick/go-example-api/pkg/domain/entity"
	"github.com/Fantamstick/go-example-api/pkg/handler"
	"github.com/Fantamstick/go-example-api/pkg/infrastracture/persistence"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	// create logger
	logger, err := zap.NewProduction()
	if err != nil {
		t.Errorf("failed to setup loggerr: %s\n", err)
	}
	defer logger.Sync()

	// load config
	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		t.Errorf("failed to load config: %s\n", err)
	}

	// init test server
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(cfg.DB.URL))
	if err != nil {
		t.Errorf("failed to create mongo db client: %s\n", err)
	}
	mongoCtx, mongoCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoCancel()

	if err := mongoClient.Connect(mongoCtx); err != nil {
		t.Errorf("failed to connect to mongo db: %s\n", err)
	}

	if err := mongoClient.Ping(mongoCtx, readpref.Primary()); err != nil {
		t.Errorf("failed to ping mongo db: %s\n", err)
	}

	mongoDB := mongoClient.Database(cfg.DB.Database)

	repositories, err := persistence.NewRepositories(mongoDB)
	assert.NoError(t, err)

	registry := handler.NewHandler(logger, repositories, "v1.0-test")
	s := NewServer(registry, &Config{Log: logger})
	testServer := httptest.NewServer(s.Mux)
	defer testServer.Close()

	t.Run("check /healthz", func(t *testing.T) {
		res, err := http.Get(testServer.URL + "/healthz")
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, http.StatusOK)
	})

	t.Run("check /version", func(t *testing.T) {
		res, err := http.Get(testServer.URL + "/version")
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, http.StatusOK)

		// read body
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		assert.NoError(t, err)
		assert.NotEmpty(t, body)

		var data interface{}
		err = json.Unmarshal(body, &data)
		assert.NoError(t, err)
		ver := data.(map[string]interface{})["version"].(string)
		assert.Equal(t, ver, "v1.0-test")
	})

	t.Run("check /user/list", func(t *testing.T) {
		res, err := http.Get(testServer.URL + "/user/list")
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, http.StatusOK)

		// read body
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		assert.NoError(t, err)
		assert.NotEmpty(t, body)

		var users []entity.User
		err = json.Unmarshal(body, &users)
		assert.NoError(t, err)
		assert.Greater(t, len(users), 0)

		assert.Contains(t, users, entity.User{
			Id:   1,
			Name: "Hoge",
		})
		assert.Contains(t, users, entity.User{
			Id:   2,
			Name: "Fuga",
		})
	})

	t.Run("check /message/1", func(t *testing.T) {
		res, err := http.Get(testServer.URL + "/message/list/1")
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, http.StatusOK)

		// read body
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		assert.NoError(t, err)
		assert.NotEmpty(t, body)

		var messages []entity.Message
		err = json.Unmarshal(body, &messages)
		assert.NoError(t, err)
		assert.Greater(t, len(messages), 0)

		assert.Contains(t, messages, entity.Message{
			Id:      1,
			UserId:  1,
			Message: "test message id 1",
		})
		assert.Contains(t, messages, entity.Message{
			Id:      2,
			UserId:  1,
			Message: "test message id 2",
		})
	})

	t.Run("check /message/2", func(t *testing.T) {
		res, err := http.Get(testServer.URL + "/message/list/2")
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, http.StatusOK)

		// read body
		body, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		assert.NoError(t, err)
		assert.NotEmpty(t, body)

		var messages []entity.Message
		err = json.Unmarshal(body, &messages)
		assert.NoError(t, err)
		assert.Greater(t, len(messages), 0)

		assert.Contains(t, messages, entity.Message{
			Id:      3,
			UserId:  2,
			Message: "test message id 3",
		})
		assert.Contains(t, messages, entity.Message{
			Id:      4,
			UserId:  2,
			Message: "test message id 4",
		})
	})
}
