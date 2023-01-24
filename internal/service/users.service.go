package service

import (
	"context"
	"errors"
	"max-inventory/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid password")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {

	U, _ := s.repo.GetUserByEmail(ctx, email)
	if U != nil {
		return ErrUserAlreadyExists
	}

	return s.repo.SaveUser(ctx, email, name, password)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	U, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if U.Password != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    U.ID,
		Email: U.Email,
		Name:  U.Name,
	}, nil
}
