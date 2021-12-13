package adopters

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Notification struct {
	PostId   string `db:"post_id"`
	Username string `db:"username"`
}

type MONGOFeedRepository struct {
	db *mongo.Collection
}

func NewMONGOFeedRepository(db *mongo.Collection) *MONGOFeedRepository {
	if db == nil {
		panic("missing db")
	}

	return &MONGOFeedRepository{db: db}
}

type sqlContextGetter interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func (m MONGOFeedRepository) CreateNotification(ctx context.Context, feed domain.Feed) error {
	_, err := m.db.InsertOne(ctx, feed)
	return err
}

func NewMongoConnection() (*mongo.Collection, error) {
	var collection *mongo.Collection
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:57017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	collection = client.Database("golang_users_db").Collection("golang_feed")

	return collection, nil
}
