package command

import (
	"context"
	"time"

	"github.com/danial2026/golang-profile-kafka/internal/profile/domain"
)

type CreateUserHandler struct {
	userRepo domain.Repository
}

func NewCreateUserHandler(userRepo domain.Repository) CreateUserHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}

	return CreateUserHandler{userRepo: userRepo}
}

func (c CreateUserHandler) Handle(ctx context.Context, users []time.Time) error {
	for _, userToUpdate := range users {
		if err := c.userRepo.CreatUser(ctx, userToUpdate); err != nil {
			return err
		}
	}

	return nil
}
