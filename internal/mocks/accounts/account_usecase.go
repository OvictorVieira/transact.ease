package accounts

import (
	"context"
	"github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	"github.com/stretchr/testify/mock"
)

type MockAccountUsecase struct {
	mock.Mock
}

func (m *MockAccountUsecase) Create(ctx context.Context, inAccount *accounts.AccountDto) (outAccount accounts.AccountDto, statusCode int, err error) {
	args := m.Called(ctx, inAccount)
	return args.Get(0).(accounts.AccountDto), args.Int(1), args.Error(2)
}

func (m *MockAccountUsecase) GetById(ctx context.Context, inAccount *accounts.AccountDto) (accounts.AccountDto, int, error) {
	args := m.Called(ctx, inAccount)
	return args.Get(0).(accounts.AccountDto), args.Int(1), args.Error(2)
}
