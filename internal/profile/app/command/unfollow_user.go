package command

import (
	"context"
	"time"

	"github.com/danial2026/golang-profile-kafka/internal/profile/domain"
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

func (c UnfollowHandler) Handle(ctx context.Context, users []time.Time) error {
	for _, userToUpdate := range users {
		if err := c.userRepo.CreatUser(ctx, userToUpdate); err != nil {
			return err
		}
	}

	return nil
}
