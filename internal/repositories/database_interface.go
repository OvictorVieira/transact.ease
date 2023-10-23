package repositories

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Database interface {
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Begin() (*sql.Tx, error)
}

type SqlxDatabase struct {
	DB *sqlx.DB
}

func (s SqlxDatabase) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	return s.DB.NamedQueryContext(ctx, query, arg)
}

func (s SqlxDatabase) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return s.DB.GetContext(ctx, dest, query, args...)
}

func (s SqlxDatabase) Begin() (*sql.Tx, error) {
	return s.DB.Begin()
}
