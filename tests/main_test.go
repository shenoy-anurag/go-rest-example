package tests

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/shenoy-anurag/go-rest-example/posts"

	"github.com/shenoy-anurag/go-rest-example/users"

	"github.com/shenoy-anurag/go-rest-example/routes"

	"github.com/shenoy-anurag/go-rest-example/configs"

	"github.com/shenoy-anurag/go-rest-example/communities"

	"go.mongodb.org/mongo-driver/mongo"
)

var router http.Handler

// customed request headers for token authorization and so on
var customHeaders = make(map[string]string, 0)

func init() {
	router = routes.SetupRouter()
	configs.LoadEnvVariables()
}

func AddHeader(key string, value string) {
	customHeaders[key] = value
}

func CreateCollections(MongoDB *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collections := []string{users.UsersCollectionName, communities.CommunitiesCollectionName}
	for i := range collections {
		MongoDB.CreateCollection(ctx, collections[i])
	}
}

func ReplaceCollectionObjects(MongoDB *mongo.Database) {
	users.UsersCollection = configs.GetCollection(MongoDB, users.UsersCollectionName)
	communities.CommunityCollection = configs.GetCollection(MongoDB, communities.CommunitiesCollectionName)
	posts.PostsCollection = configs.GetCollection(MongoDB, posts.PostsCollectionName)
}

func TestMain(m *testing.M) {
	log.Println("Database setup for test")
	// user and password will need to match running postgres instance
	configs.MongoDB = configs.MongoClient.Database(os.Getenv("TEST_DB_NAME"))
	// Check if DB is connected
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := configs.MongoClient.Ping(ctx, nil)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	CreateCollections(configs.MongoDB)
	ReplaceCollectionObjects(configs.MongoDB)

	exitVal := m.Run()

	log.Println("Database dropped after test")

	os.Exit(exitVal)
}
