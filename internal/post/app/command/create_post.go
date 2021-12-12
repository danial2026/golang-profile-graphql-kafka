package command

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/post/domain"
)

type CreatePostHandler struct {
	postRepo domain.Repository
}

func NewCreatePost(postRepo domain.Repository) CreatePostHandler {
	if postRepo == nil {
		panic("userRepo is nil")
	}

	return CreatePostHandler{postRepo: postRepo}
}

func (c CreatePostHandler) Handle(ctx context.Context, user domain.Post) error {
	if err := c.postRepo.CreatPost(ctx, user); err != nil {
		return err
	}

	return nil
}
