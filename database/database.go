package database

import (
	"context"
	"fmt"
	"github.com/go-inventory/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota", s.DB.Host, s.DB.User, s.DB.Password, s.DB.Name, s.DB.Port)
	return sqlx.ConnectContext(ctx, "postgres", dsn)

}
