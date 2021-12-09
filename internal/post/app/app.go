package app

import (
	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreatePost command.CreatePostHandler
}

type Queries struct {
}
