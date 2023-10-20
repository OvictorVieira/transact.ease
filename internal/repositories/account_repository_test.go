package repositories

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
	"testing"

	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	Mocks "github.com/OvictorVieira/transact.ease/internal/mocks/accounts"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndReturnData(t *testing.T) {
	mockDB := new(Mocks.MockDatabase)
	repo := NewAccountRepository(mockDB)

	accountDto := &Domain.AccountDto{
		DocumentNumber: "123456789",
	}

	mockDB.On("NamedQueryContext", mock.Anything, mock.Anything, mock.Anything).Return(&sqlx.Rows{}, nil)

	err := repo.Create(context.TODO(), accountDto)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestCreateAndReturnError(t *testing.T) {
	mockDB := new(Mocks.MockDatabase)
	repo := NewAccountRepository(mockDB)

	accountDto := &Domain.AccountDto{
		DocumentNumber: "",
	}

	mockDB.On("NamedQueryContext", mock.Anything, mock.Anything, mock.Anything).Return(&sqlx.Rows{}, errors.New("error"))

	err := repo.Create(context.TODO(), accountDto)

	assert.Error(t, err)
	mockDB.AssertExpectations(t)
}

func TestGetByDocumentNumberAndReturnData(t *testing.T) {
	mockDB := new(Mocks.MockDatabase)
	repo := NewAccountRepository(mockDB)

	accountDto := &Domain.AccountDto{
		DocumentNumber: "123456789",
	}

	expectedAccount := Domain.AccountDto{
		DocumentNumber: "123456789",
	}

	mockDB.On("GetContext", mock.Anything, mock.Anything, mock.Anything, accountDto.DocumentNumber).Return(nil)

	actualAccount, err := repo.GetByDocumentNumber(context.TODO(), accountDto)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, actualAccount)
	mockDB.AssertExpectations(t)
}
