package command

import (
	"context"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/feed/domain"
)

type CreateNotificationHandler struct {
	feedRepo domain.Repository
}

func NewCreateNotification(feedRepo domain.Repository) CreateNotificationHandler {
	if feedRepo == nil {
		panic("userRepo is nil")
	}

	return CreateNotificationHandler{feedRepo: feedRepo}
}

func (c CreateNotificationHandler) Handle(ctx context.Context, feed domain.Feed) error {
	if err := c.feedRepo.CreateNotification(ctx, feed); err != nil {
		return err
	}

	return nil
}
