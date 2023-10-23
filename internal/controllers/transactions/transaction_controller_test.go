package transactions

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/OvictorVieira/transact.ease/internal/constants"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
	Request "github.com/OvictorVieira/transact.ease/internal/dto/requests"
	Mocks "github.com/OvictorVieira/transact.ease/internal/mocks/transactions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestProcessTransactionCreationWithSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUsecase := new(Mocks.MockTransactionUsecase)
	controller := NewTransactionController(mockUsecase)

	accountDto := Domain.TransactionDto{
		ID:              1,
		AccountId:       1,
		OperationTypeId: 1,
		Amount:          52,
		CreatedAt:       time.Now().In(constants.UTC),
		UpdatedAt:       time.Now().In(constants.UTC),
	}

	mockUsecase.On("Create", mock.Anything, mock.Anything).Return(accountDto, http.StatusCreated, nil)

	accountCreationRequest := Request.TransactionCreationRequest{
		AccountId:       1,
		OperationTypeId: 1,
		Amount:          52,
	}

	requestBody, _ := json.Marshal(accountCreationRequest)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(requestBody))

	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/transactions", controller.ProcessTransactionCreation)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestProcessTransactionCreationWithErrorFieldRequired(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUsecase := new(Mocks.MockTransactionUsecase)
	controller := NewTransactionController(mockUsecase)

	mockUsecase.On("Create", mock.Anything, mock.Anything).Return(Domain.TransactionDto{}, http.StatusBadRequest, errors.New("error message"))

	accountCreationRequest := Request.TransactionCreationRequest{
		AccountId:       1,
		OperationTypeId: 1,
	}

	requestBody, _ := json.Marshal(accountCreationRequest)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(requestBody))

	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/transactions", controller.ProcessTransactionCreation)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
