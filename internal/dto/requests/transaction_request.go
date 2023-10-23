package requests

import (
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
)

type TransactionCreationRequest struct {
	AccountId       int     `json:"account_id" validate:"required"`
	OperationTypeId int     `json:"operation_type_id" validate:"required"`
	Amount          float32 `json:"amount" validate:"required"`
}

func (a *TransactionCreationRequest) ToTransactionDto() *Domain.TransactionDto {
	return &Domain.TransactionDto{
		AccountId:       a.AccountId,
		OperationTypeId: a.OperationTypeId,
		Amount:          a.Amount,
	}
}
