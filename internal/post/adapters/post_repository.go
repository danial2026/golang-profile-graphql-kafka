package adopters

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Post struct {
// 	ID          string `db:"id"`
// 	ProfileId   string `db:"profile_id"`
// 	Body        string `db:"body"`
// 	CreatedTime string `db:"created_time"`
// }

type Liked struct {
	PostId   string `db:"post_id"`
	Username string `db:"username"`
}

type MONGOPostRepository struct {
	db *mongo.Collection
}

func NewMONGOPostRepository(db *mongo.Collection) *MONGOPostRepository {
	if db == nil {
		panic("missing db")
	}

	return &MONGOPostRepository{db: db}
}

type sqlContextGetter interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func (m MONGOPostRepository) CreatPost(ctx context.Context, post domain.Post) error {
	_, err := m.db.InsertOne(ctx, post)
	return err
}

func (m MONGOPostRepository) Like(ctx context.Context, username string, postId string) error {
	_, err := m.db.InsertOne(ctx,
		Liked{
			PostId:   postId,
			Username: username,
		},
	)
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

	collection = client.Database("golang_users_db").Collection("golang_post")

	return collection, nil
}
