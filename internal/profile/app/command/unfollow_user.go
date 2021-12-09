package command

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/domain"
)

type UnfollowHandler struct {
	userRepo domain.Repository
}

func NewUnfollowHandler(userRepo domain.Repository) UnfollowHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}

	return UnfollowHandler{userRepo: userRepo}
}

func (c UnfollowHandler) Handle(ctx context.Context, username1 string, username2 string) error {
	if err := c.userRepo.Unfollow(ctx, username1, username2); err != nil {
		return err
	}

	return nil
}
