package command

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/profile/domain"
)

type CreateAccountHandler struct {
	userRepo domain.Repository
}

func NewCreateAccountHandler(userRepo domain.Repository) CreateAccountHandler {
	if userRepo == nil {
		panic("userRepo is nil")
	}

	return CreateAccountHandler{userRepo: userRepo}
}

func (c CreateAccountHandler) Handle(ctx context.Context, user domain.User) error {
	if err := c.userRepo.CreatUser(ctx, user); err != nil {
		return err
	}

	return nil
}
