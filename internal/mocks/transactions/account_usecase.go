package transactions

import (
	"context"
	"github.com/OvictorVieira/transact.ease/internal/domains/transactions"
	"github.com/stretchr/testify/mock"
)

type MockTransactionUsecase struct {
	mock.Mock
}

func (m *MockTransactionUsecase) Create(ctx context.Context, inTransaction *transactions.TransactionDto) (outTransaction transactions.TransactionDto, statusCode int, err error) {
	args := m.Called(ctx, inTransaction)
	return args.Get(0).(transactions.TransactionDto), args.Int(1), args.Error(2)
}

func (m *MockTransactionUsecase) GetById(ctx context.Context, inTransaction *transactions.TransactionDto) (transactions.TransactionDto, int, error) {
	args := m.Called(ctx, inTransaction)
	return args.Get(0).(transactions.TransactionDto), args.Int(1), args.Error(2)
}
