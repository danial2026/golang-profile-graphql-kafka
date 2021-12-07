package command

import (
	"context"
	"os/user"
	"time"

	"github.com/danial2026/golang-profile-kafk/internal/profile/domain"
)

type CreateUserHandler struct {
	userRepo user.Repository
}

func NewCreateUserHandler(userRepo user.Repository) CreateUserHandler {
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
