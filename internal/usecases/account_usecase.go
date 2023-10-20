package usecases

import (
	"context"
	"github.com/OvictorVieira/transact.ease/internal/constants"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	LOGGER "github.com/OvictorVieira/transact.ease/pkg/logger"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type accountUsecase struct {
	repo Domain.AccountRepository
}

func NewUserService(repo Domain.AccountRepository) Domain.AccountUsecase {
	return &accountUsecase{
		repo: repo,
	}
}

func (a accountUsecase) Create(ctx context.Context, inAccount *Domain.AccountDto) (outAccount Domain.AccountDto, statusCode int, err error) {
	inAccount.CreatedAt = time.Now().In(constants.UTC)
	inAccount.UpdatedAt = time.Now().In(constants.UTC)

	LOGGER.Info("trying to create an account", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})

	err = a.repo.Create(ctx, inAccount)
	if err != nil {
		LOGGER.Error("error when try to create account", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})

		if strings.Contains(err.Error(), "duplicate key value") {
			return Domain.AccountDto{}, http.StatusBadRequest, constants.ErrDuplicatedDocumentNumber
		} else {
			LOGGER.Error(err.Error(), logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
			return Domain.AccountDto{}, http.StatusInternalServerError, constants.ErrCreateAccount
		}
	}

	LOGGER.Info("trying to find account after create", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})

	outAccount, err = a.repo.GetByDocumentNumber(ctx, inAccount)
	if err != nil {
		LOGGER.Error("error when try to find account after create", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
		return Domain.AccountDto{}, http.StatusInternalServerError, constants.ErrAccountNotFound
	}

	LOGGER.Info("account created with success", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})

	return outAccount, http.StatusCreated, nil
}

func (a accountUsecase) GetById(ctx context.Context, inAccount *Domain.AccountDto) (outAccount Domain.AccountDto, statusCode int, err error) {
	LOGGER.InfoF("trying to get an account by id: ", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow}, inAccount.ID)

	outAccount, err = a.repo.GetById(ctx, inAccount)
	if err != nil {
		LOGGER.Error("error when try to find account with id provided", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})
		return Domain.AccountDto{}, http.StatusNotFound, constants.ErrAccountNotFound
	}

	LOGGER.Info("account found", logrus.Fields{constants.LoggerCategory: constants.LoggerCategorySystemFlow})

	return outAccount, http.StatusOK, nil
}
