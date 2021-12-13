package adopters

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type Notification struct {
	PostId   string `db:"post_id"`
	Username string `db:"username"`
}

type MongoFeedRepository struct {
	db *mongo.Client
}

func NewMongoFeedRepository(cli *mongo.Client) MongoFeedRepository {
	return MongoFeedRepository{
		db: cli,
	}
}

type sqlContextGetter interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func (m MongoFeedRepository) feedCollection() *mongo.Collection {
	return m.db.Database("golang_users_db").Collection("golang_feed")
}

func (m MongoFeedRepository) CreateNotification(ctx context.Context, feed domain.Feed) error {
	_, err := m.feedCollection().InsertOne(ctx, feed)
	return err
}