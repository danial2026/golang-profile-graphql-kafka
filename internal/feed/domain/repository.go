package domain

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	username string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("user '%s' not found", e.username)
}

type Repository interface {
	CreateNotification(
		ctx context.Context,
		feed Feed) error
}
