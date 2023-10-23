package transactions

import (
	Controller "github.com/OvictorVieira/transact.ease/internal/controllers/transactions"
	Repository "github.com/OvictorVieira/transact.ease/internal/repositories"
	Usecase "github.com/OvictorVieira/transact.ease/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type TransactionRoutes struct {
	Controller Controller.TransactionController
	router     *gin.RouterGroup
	db         *sqlx.DB
}

func NewTransactionsRoute(router *gin.RouterGroup, db *sqlx.DB) *TransactionRoutes {
	TransactionRepository := Repository.NewTransactionRepository(db)
	TransactionUsecase := Usecase.NewTransactionUsecase(TransactionRepository)
	TransactionController := Controller.NewTransactionController(TransactionUsecase)

	return &TransactionRoutes{
		Controller: TransactionController,
		router:     router,
		db:         db,
	}
}

func (ar *TransactionRoutes) Routes() {
	{
		router := ar.router.Group("/transactions")
		router.POST("/", ar.Controller.ProcessTransactionCreation)
	}
}
