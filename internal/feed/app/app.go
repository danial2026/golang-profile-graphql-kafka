package app

import (
	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateNotification command.CreateNotificationHandler
}

type Queries struct {
}
