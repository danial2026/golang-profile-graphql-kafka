package command

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/domain"
)

type LikePostHandler struct {
	postRepo domain.Repository
}

func LikePost(postRepo domain.Repository) LikePostHandler {
	if postRepo == nil {
		panic("postRepo is nil")
	}

	return LikePostHandler{postRepo: postRepo}
}

func (c LikePostHandler) Handle(ctx context.Context, username string, postId string) error {
	if err := c.postRepo.Like(ctx, username, postId); err != nil {
		return err
	}

	return nil
}
