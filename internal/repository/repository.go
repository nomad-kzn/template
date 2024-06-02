package repository

import (
	"context"
	"database/sql"

	"github.com/nomad-kzn/template/internal/interfaces"
	"github.com/nomad-kzn/template/internal/models"
)

var _ interfaces.UserRepository = (*RepoImpl)(nil)

type RepoImpl struct {
	db *sql.DB
}

func NewUserRepoImpl(db *sql.DB) *RepoImpl {
	return &RepoImpl{db: db}
}

func (r *RepoImpl) GetUser(ctx context.Context, id string) (*models.UserEntity, error) {
	return &models.UserEntity{}, nil
}