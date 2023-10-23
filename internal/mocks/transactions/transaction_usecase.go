package transactions

import (
	"context"
	"github.com/OvictorVieira/transact.ease/internal/domains/transactions"
	"github.com/stretchr/testify/mock"
)

type MockTransactionUsecase struct {
	mock.Mock
}

func (m *MockTransactionUsecase) Create(inTransaction *transactions.TransactionDto) (ouTransaction *transactions.TransactionDto, statusCode int, err error) {
	args := m.Called(inTransaction)

	transactionDtoValue := args.Get(0).(transactions.TransactionDto)
	statusCode = args.Int(1)
	err = args.Error(2)

	return &transactionDtoValue, statusCode, err
}

func (m *MockTransactionUsecase) GetById(ctx context.Context, inTransaction *transactions.TransactionDto) (transactions.TransactionDto, int, error) {
	args := m.Called(ctx, inTransaction)
	return args.Get(0).(transactions.TransactionDto), args.Int(1), args.Error(2)
}
