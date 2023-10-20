package accounts

import (
	"bytes"
	"encoding/json"
	"github.com/OvictorVieira/transact.ease/internal/constants"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	Mocks "github.com/OvictorVieira/transact.ease/internal/mocks/accounts"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	Request "github.com/OvictorVieira/transact.ease/internal/dto/requests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProcessAccountCreationWithSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUsecase := new(Mocks.MockAccountUsecase)
	controller := NewAccountController(mockUsecase)

	accountDto := Domain.AccountDto{
		ID:             1,
		DocumentNumber: "123456789",
		CreatedAt:      time.Now().In(constants.UTC),
		UpdatedAt:      time.Now().In(constants.UTC),
	}

	mockUsecase.On("Create", mock.Anything, mock.Anything).Return(accountDto, http.StatusCreated, nil)

	accountCreationRequest := Request.AccountCreationRequest{
		DocumentNumber: "123456789",
	}

	requestBody, _ := json.Marshal(accountCreationRequest)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(requestBody))

	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/accounts", controller.ProcessAccountCreation)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestProcessAccountCreationWithoutDocumentNumberWithBadRequestError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUsecase := new(Mocks.MockAccountUsecase)
	controller := NewAccountController(mockUsecase)

	accountDto := Domain.AccountDto{}

	mockUsecase.On("Create", mock.Anything, mock.Anything).Return(accountDto, http.StatusBadRequest, nil)

	accountCreationRequest := Request.AccountCreationRequest{
		DocumentNumber: "",
	}

	requestBody, _ := json.Marshal(accountCreationRequest)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(requestBody))

	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/accounts", controller.ProcessAccountCreation)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestProcessAccountCreationWithoutPayloadWithBadRequestError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUsecase := new(Mocks.MockAccountUsecase)
	controller := NewAccountController(mockUsecase)

	accountDto := Domain.AccountDto{}

	mockUsecase.On("Create", mock.Anything, mock.Anything).Return(accountDto, http.StatusBadRequest, nil)

	accountCreationRequest := Request.AccountCreationRequest{
		DocumentNumber: "",
	}

	requestBody, _ := json.Marshal(accountCreationRequest)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(requestBody))

	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/accounts", controller.ProcessAccountCreation)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestProcessAccountCreationWithInternalServerError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUsecase := new(Mocks.MockAccountUsecase)
	controller := NewAccountController(mockUsecase)

	accountDto := Domain.AccountDto{}

	mockUsecase.On("Create", mock.Anything, mock.Anything).Return(accountDto, http.StatusInternalServerError, nil)

	accountCreationRequest := Request.AccountCreationRequest{
		DocumentNumber: "123456789",
	}

	requestBody, _ := json.Marshal(accountCreationRequest)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(requestBody))

	w := httptest.NewRecorder()

	router := gin.Default()
	router.POST("/accounts", controller.ProcessAccountCreation)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
