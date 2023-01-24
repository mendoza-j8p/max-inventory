package service

import (
	"context"
	"max-inventory/internal/models"
	"max-inventory/internal/repository"
)


type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User,error)
}

type serv struct {
	repo repository.Respository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo, 
	}
}