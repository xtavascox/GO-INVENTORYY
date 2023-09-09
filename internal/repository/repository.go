package repository

import (
	"context"
	"github.com/go-inventory/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository is an interface that defines the methods that any
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, password, name string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
