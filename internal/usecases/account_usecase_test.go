package usecases

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"

	"github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	Mocks "github.com/OvictorVieira/transact.ease/internal/mocks/accounts"
	"github.com/stretchr/testify/assert"
)

func TestCreateWithSuccess(t *testing.T) {
	mockRepo := new(Mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)

	accountDto := &accounts.AccountDto{
		DocumentNumber: "123456789",
	}

	mockRepo.On("Create", mock.Anything, accountDto).Return(nil).Once()
	mockRepo.On("GetByDocumentNumber", mock.Anything, accountDto).Return(*accountDto, nil).Once()

	outAccount, statusCode, err := usecase.Create(context.TODO(), accountDto)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, statusCode)
	assert.Equal(t, accountDto.ID, outAccount.ID)

	mockRepo.AssertExpectations(t)
}

func TestCreateFailsWithDuplicatedKey(t *testing.T) {
	mockRepo := new(Mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)

	accountDto := &accounts.AccountDto{
		DocumentNumber: "123456789",
	}

	mockRepo.On("Create", mock.Anything, accountDto).Return(errors.New("duplicate key value")).Once()

	_, statusCode, err := usecase.Create(context.TODO(), accountDto)

	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, statusCode)

	mockRepo.AssertExpectations(t)
}

func TestCreateFailsWithGenericError(t *testing.T) {
	mockRepo := new(Mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)

	accountDto := &accounts.AccountDto{
		DocumentNumber: "123456789",
	}

	mockRepo.On("Create", mock.Anything, accountDto).Return(errors.New("db off")).Once()

	_, statusCode, err := usecase.Create(context.TODO(), accountDto)

	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, statusCode)

	mockRepo.AssertExpectations(t)
}

func TestGetByIdWithSuccess(t *testing.T) {
	mockRepo := new(Mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)

	accountDto := &accounts.AccountDto{
		ID: 1,
	}

	mockRepo.On("GetById", mock.Anything, accountDto).Return(*accountDto, nil)

	_, statusCode, err := usecase.GetById(context.TODO(), accountDto)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

	mockRepo.AssertExpectations(t)
}

func TestCreateFailsWithNotFound(t *testing.T) {
	mockRepo := new(Mocks.MockAccountRepository)
	usecase := NewAccountUsecase(mockRepo)

	accountDto := &accounts.AccountDto{
		ID: -1,
	}

	mockRepo.On("GetById", mock.Anything, accountDto).Return(accounts.AccountDto{}, errors.New("not found")).Once()

	_, statusCode, err := usecase.GetById(context.TODO(), accountDto)

	assert.Error(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)

	mockRepo.AssertExpectations(t)
}
