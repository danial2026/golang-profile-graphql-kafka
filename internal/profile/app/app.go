package app

import (
	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateAccount command.CreateAccountHandler

	Follow   command.FollowHandler
	Unfollow command.UnfollowHandler
}

type Queries struct {
}
