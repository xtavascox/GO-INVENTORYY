package services

import (
	"context"
	"errors"
	"github.com/go-inventory/encryption"
	"github.com/go-inventory/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *serv) RegisterUser(ctx context.Context, email string, password string, name string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}
	ec, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}
	pass := encryption.ToBase64(ec)
	return s.repo.SaveUser(ctx, email, pass, name)
}

func (s *serv) LoginUser(ctx context.Context, email string, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	dc, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}
	decryptedPassword, err := encryption.Decrypt(dc)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
