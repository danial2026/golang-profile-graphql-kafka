package domain

import (
	"context"
	"fmt"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/domain"
)

type NotFoundError struct {
	username string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("user '%s' not found", e.username)
}

type Repository interface {
	CreatPost(
		ctx context.Context,
		post domain.Post) error

	Like(
		ctx context.Context,
		username string,
		postId string) error
}
