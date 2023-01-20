package repository

import (
	"context"
	"database/sql"
	"max-inventory/internal/entity"
)

type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}
