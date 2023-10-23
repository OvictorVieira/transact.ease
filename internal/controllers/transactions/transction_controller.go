package transactions

import (
	"github.com/OvictorVieira/transact.ease/internal/constants"
	Controllers "github.com/OvictorVieira/transact.ease/internal/controllers"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
	Request "github.com/OvictorVieira/transact.ease/internal/dto/requests"
	"github.com/OvictorVieira/transact.ease/internal/dto/responses"
	"github.com/OvictorVieira/transact.ease/pkg/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionController struct {
	usecase Domain.TransactionUsecase
}

func NewTransactionController(usecase Domain.TransactionUsecase) TransactionController {
	return TransactionController{
		usecase: usecase,
	}
}

func (ac TransactionController) ProcessTransactionCreation(ctx *gin.Context) {
	var transactionCreationRequest Request.TransactionCreationRequest

	if err := ctx.ShouldBindJSON(&transactionCreationRequest); err != nil {
		Controllers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(transactionCreationRequest); err != nil {
		Controllers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	transactionDto, statusCode, err := ac.usecase.Create(ctx.Request.Context(), transactionCreationRequest.ToTransactionDto())
	if err != nil {
		Controllers.NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	Controllers.NewSuccessResponse(ctx, statusCode, constants.TransactionCreatedWithSuccess, responses.FromTransactionDto(transactionDto))
}
