package adopters

import (
	"context"
	"time"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	Email     string    `db:"email"`
	Username  string    `db:"username"`
	Following []User    `db:"following"`
	Followers []User    `db:"followers"`
}

type MONGOUserRepository struct {
	db          *mongo.Collection
}

func NewMONGOUserRepository(db *mongo.Collection) *MONGOUserRepository {
	if db == nil {
		panic("missing db")
	}

	return &MONGOUserRepository{db: db}
}

type sqlContextGetter interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func (m MONGOUserRepository) CreatUser(ctx context.Context, user domain.User) error {
	_, err := m.db.InsertOne(ctx, user)
	return err
}

func (m MONGOUserRepository) Follow(ctx context.Context, username1 string, username2 string) error {
	_, err := m.db.Update(ctx,
		bson.M{"username": username1},
		bson.M{
			{
				"$push": bson.M{"Following": bson.M{username2}}
			},
		}
		)
	return err
}

func (m MONGOUserRepository) Unfollow(ctx context.Context, username1 string, username2 string) error {
	_, err := m.db.Update(ctx,
		bson.M{"username": username1},
		bson.M{
			{
				"$pull": bson.M{"Following": bson.M{"$in": bson.M{username2}}}
			},
		}
		)
	return err
}

func NewMongoConnection() (*mongo.Collection, error) {
	var collection *mongo.Collection
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:37017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	collection = client.Database("golang_users_db").Collection("golang_users")

	return collection, nil
}
