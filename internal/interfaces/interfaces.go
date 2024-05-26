package interfaces

import (
	"context"
	"github.com/nomad-kzn/template/internal/models"
)

type (
	Usecase interface {
		GetUser(ctx context.Context, id string) (*models.User, error)
	}
)
