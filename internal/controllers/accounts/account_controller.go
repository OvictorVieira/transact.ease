package accounts

import (
	"github.com/OvictorVieira/transact.ease/internal/constants"
	Controllers "github.com/OvictorVieira/transact.ease/internal/controllers"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	Request "github.com/OvictorVieira/transact.ease/internal/dto/requests"
	"github.com/OvictorVieira/transact.ease/internal/dto/responses"
	"github.com/OvictorVieira/transact.ease/pkg/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AccountController struct {
	usecase Domain.AccountUsecase
}

func NewAccountController(usecase Domain.AccountUsecase) AccountController {
	return AccountController{
		usecase: usecase,
	}
}

func (ac AccountController) ProcessAccountCreation(ctx *gin.Context) {
	var accountCreationRequest Request.AccountCreationRequest

	if err := ctx.ShouldBindJSON(&accountCreationRequest); err != nil {
		Controllers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validators.ValidatePayloads(accountCreationRequest); err != nil {
		Controllers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	accountDto, statusCode, err := ac.usecase.Create(ctx.Request.Context(), accountCreationRequest.ToAccountDto())
	if err != nil {
		Controllers.NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	Controllers.NewSuccessResponse(ctx, statusCode, constants.AccountCreatedWithSuccess, responses.FromAccountDto(accountDto))
}

func (ac AccountController) GetById(ctx *gin.Context) {
	accountId := ctx.Param("accountId")

	convertedAccountId, err := strconv.Atoi(accountId)

	if err != nil {
		Controllers.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	accountDto, statusCode, err := ac.usecase.GetById(ctx.Request.Context(), Request.BuildAccountDtoToFind(convertedAccountId))
	if err != nil {
		Controllers.NewErrorResponse(ctx, statusCode, err.Error())
		return
	}

	Controllers.NewSuccessResponse(ctx, statusCode, constants.AccountCreatedWithSuccess, responses.FromAccountDto(accountDto))
}
