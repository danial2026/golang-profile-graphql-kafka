package command

import (
	"context"

	"github.com/danial2026/golang-profile-kafka/internal/profile/domain"
)

type FollowHandler struct {
	userRepo domain.Repository
}

func NewFollowHandler(userRepo domain.Repository) FollowHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}

	return FollowHandler{userRepo: userRepo}
}

func (c FollowHandler) Handle(ctx context.Context, username1 string, username2 string) error {
	if err := c.userRepo.Follow(ctx, username1, username2); err != nil {
		return err
	}

	return nil
}
