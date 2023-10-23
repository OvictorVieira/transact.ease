package repositories

import (
	"context"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
	Mocks "github.com/OvictorVieira/transact.ease/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTransactionRepositoryGetByAccountIdAndOperationTypeIdWithSuccess(t *testing.T) {
	mockDB := new(Mocks.MockDatabase)
	repo := NewTransactionRepository(mockDB)

	transactionDto := &Domain.TransactionDto{
		AccountId:       1,
		OperationTypeId: 1,
	}

	expectedTransaction := Domain.TransactionDto{
		AccountId:       1,
		OperationTypeId: 1,
	}

	mockDB.On("GetContext", mock.Anything, mock.Anything, mock.Anything, transactionDto.AccountId, transactionDto.OperationTypeId).Return(nil)

	actualTransaction, err := repo.GetByAccountIdAndOperationTypeId(context.TODO(), transactionDto)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransaction, actualTransaction)
	mockDB.AssertExpectations(t)
}
