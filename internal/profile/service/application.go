package service

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/adapters"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/app"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/app/command"

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

	collection = client.Database("golang_users_db").Collection("golang_users")

	usersRepository := adapters.NewMONGOUserRepository(collection)

	return app.Application{
		Commands: app.Commands{
			CreateUser: command.NewCreateUserHandler(usersRepository),
			Follow:     command.NewFollowHandler(usersRepository),
			Unfollow:   command.NewUnfollowHandler(usersRepository),
		},
		Queries: app.Queries{},
	}
}
