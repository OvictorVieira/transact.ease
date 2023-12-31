package accounts

import (
	Controller "github.com/OvictorVieira/transact.ease/internal/controllers/accounts"
	Repository "github.com/OvictorVieira/transact.ease/internal/repositories"
	Usecase "github.com/OvictorVieira/transact.ease/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type AccountRoutes struct {
	Controller Controller.AccountController
	router     *gin.RouterGroup
	db         *sqlx.DB
}

func NewAccountsRoute(router *gin.RouterGroup, db *sqlx.DB) *AccountRoutes {
	accountRepository := Repository.NewAccountRepository(db)
	accountUsecase := Usecase.NewAccountUsecase(accountRepository)
	accountController := Controller.NewAccountController(accountUsecase)

	return &AccountRoutes{
		Controller: accountController,
		router:     router,
		db:         db,
	}
}

func (ar *AccountRoutes) Routes() {
	{
		router := ar.router.Group("/accounts")
		router.POST("/", ar.Controller.ProcessAccountCreation)
		router.GET("/:accountId", ar.Controller.GetById)
	}
}
