package domain

import (
	"context"
	"fmt"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/domain"
)

type NotFoundError struct {
	username string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("user '%s' not found", e.username)
}

type Repository interface {
	CreatUser(
		ctx context.Context,
		user domain.User) error

	Follow(
		ctx context.Context,
		username1 string, // currentUser
		username2 string) error // followThisUser

	Unfollow(
		ctx context.Context,
		username1 string, // currentUser
		username2 string) error // followThisUser
}
