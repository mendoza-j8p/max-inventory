package repository

import (
	"context"
	"max-inventory/internal/entity"
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	return nil
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error)  {
	return nil, nil
}
