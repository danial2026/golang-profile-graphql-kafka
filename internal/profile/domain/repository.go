package domain

import (
	"context"
	"fmt"
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
		user User) error

	Follow(
		ctx context.Context,
		username1 string, // currentUser
		username2 string) error // followThisUser

	Unfollow(
		ctx context.Context,
		username1 string, // currentUser
		username2 string) error // followThisUser
}
