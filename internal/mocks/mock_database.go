package mocks

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	args := m.Called(ctx, query, arg)
	return args.Get(0).(*sqlx.Rows), args.Error(1)
}

func (m *MockDatabase) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	argList := make([]interface{}, len(args)+3)
	argList[0] = ctx
	argList[1] = dest
	argList[2] = query
	copy(argList[3:], args)
	args = m.Called(argList...)

	return nil
}

func (m *MockDatabase) Begin() (*sql.Tx, error) {
	args := m.Called()
	return args.Get(0).(*sql.Tx), args.Error(1)
}
