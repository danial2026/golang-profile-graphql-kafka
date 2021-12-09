package service

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/adapters"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/app"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/app/command"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewApplication(ctx context.Context) app.Application {
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

	collection = client.Database("golang_users_db").Collection("golang_post")

	postRepository := adapters.NewMONGOUserRepository(collection)

	return app.Application{
		Commands: app.Commands{
			CreatePost: command.NewCreatePostHandler(postRepository),
		},
		Queries: app.Queries{},
	}
}
