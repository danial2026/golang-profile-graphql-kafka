package command

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/domain"
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

func (c CreateUserHandler) Handle(ctx context.Context, user domain.User) error {
	if err := c.userRepo.CreatUser(ctx, user); err != nil {
		return err
	}

	return nil
}
