package services

import (
	"context"
	"github.com/go-inventory/internal/models"
	"github.com/go-inventory/internal/repository"
)

// Service is an interface that defines the methods that any
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email string, password string, name string) error
	LoginUser(ctx context.Context, email string, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
