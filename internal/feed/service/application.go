package service

import (
	"context"

	adapters "github.com/danial2026/golang-profile-graphql-kafka/internal/feed/adapters"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/app"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/app/command"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewApplication(ctx context.Context) app.Application {
	var collection *mongo.Collection
	// var ctx = context.TODO()

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

	feedRepository := adapters.NewMONGOFeedRepository(collection)

	return app.Application{
		Commands: app.Commands{
			CreateNotification: command.NewCreateNotification(feedRepository),
		},
		Queries: app.Queries{},
	}
}
