package repository

import (
	"context"
	"github.com/go-inventory/internal/entity"
)

const (
	querySaveUser = `
		INSERT INTO users (email, password,name)
		VALUES ($1, $2,$3);
`
	queryGetUserByEmail = `
	SELECT id, email, password, name
	FROM users WHERE email = $1;
	`
)

func (r *repo) SaveUser(ctx context.Context, email string, password string, name string) error {
	_, err := r.db.ExecContext(ctx, querySaveUser, email, password, name)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	u := &entity.User{}
	err := r.db.GetContext(ctx, u, queryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return u, nil
}
