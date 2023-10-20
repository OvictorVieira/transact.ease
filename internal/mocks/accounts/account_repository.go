package accounts

import (
	"context"
	"github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (m *MockAccountRepository) Create(ctx context.Context, inAccount *accounts.AccountDto) error {
	args := m.Called(ctx, inAccount)
	return args.Error(0)
}

func (m *MockAccountRepository) GetByDocumentNumber(ctx context.Context, inAccount *accounts.AccountDto) (accounts.AccountDto, error) {
	args := m.Called(ctx, inAccount)
	return args.Get(0).(accounts.AccountDto), args.Error(1)
}

func (m *MockAccountRepository) GetById(ctx context.Context, inAccount *accounts.AccountDto) (outAccount accounts.AccountDto, err error) {
	args := m.Called(ctx, inAccount)
	return args.Get(0).(accounts.AccountDto), args.Error(1)
}
