package usecases

import (
	"github.com/OvictorVieira/transact.ease/internal/constants"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
	LOGGER "github.com/OvictorVieira/transact.ease/pkg/logger"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type transactionUsecase struct {
	repo Domain.TransactionRepository
}

func NewTransactionUsecase(repo Domain.TransactionRepository) Domain.TransactionUsecase {
	return &transactionUsecase{
		repo: repo,
	}
}

func (t transactionUsecase) Create(inTransaction *Domain.TransactionDto) (ouTransaction *Domain.TransactionDto, statusCode int, err error) {
	inTransaction.EventDate = time.Now().In(constants.UTC)
	inTransaction.CreatedAt = time.Now().In(constants.UTC)
	inTransaction.UpdatedAt = time.Now().In(constants.UTC)

	LOGGER.Info("trying to create a transaction", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
	createdTransactionId, err := t.repo.Create(inTransaction)
	if err != nil {
		LOGGER.Error("error when try to create a transaction", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})

		LOGGER.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
		return &Domain.TransactionDto{}, http.StatusInternalServerError, constants.ErrCreateTransaction
	}

	inTransaction.ID = createdTransactionId

	LOGGER.Info("transaction created with success", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})

	return inTransaction, http.StatusCreated, nil
}
