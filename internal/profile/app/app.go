package app

import (
	"github.com/danial2026/golang-profile-kafka/internal/profile/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser   command.CreateUserHandler

	Follow   command.FollowHandler
	Unfollow command.UnfollowHandler
}

type Queries struct {
}
