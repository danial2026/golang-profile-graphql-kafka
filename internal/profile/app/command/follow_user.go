package command

import (
	"context"
	"time"

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

func (c FollowHandler) Handle(ctx context.Context, users []time.Time) error {
	for _, userToUpdate := range users {
		if err := c.userRepo.CreatUser(ctx, userToUpdate); err != nil {
			return err
		}
	}

	return nil
}
