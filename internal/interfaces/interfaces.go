package interfaces

import (
	"context"
	"github.com/nomad-kzn/template/internal/models"
)

type (
	Usecase interface {
		GetUser(ctx context.Context, id string) (*models.User, error)
	}

	Repository	interface {
		UserRepository
	}

	UserRepository interface {
		GetUser(ctx context.Context, id string) (*models.UserEntity, error)
	}

	Srvc interface {
		GetUsersStats(ctx context.Context) ([]models.UserStat, error)
	}
)
