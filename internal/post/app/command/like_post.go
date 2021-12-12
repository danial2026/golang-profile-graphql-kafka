package command

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/domain"
)

type LikePostHandler struct {
	userRepo domain.Repository
}

func LikePost(userRepo domain.Repository) LikePostHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}

	return LikePostHandler{userRepo: userRepo}
}

func (c LikePostHandler) Handle(ctx context.Context, username string, postId string) error {
	if err := c.userRepo.Like(ctx, username, postId); err != nil {
		return err
	}

	return nil
}
