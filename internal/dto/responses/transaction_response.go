package responses

import (
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
)

type TransactionResponse struct {
	Id int `json:"id"`
}

func (ar *TransactionResponse) ToTransactionDto() Domain.TransactionDto {
	return Domain.TransactionDto{
		ID: ar.Id,
	}
}

func FromTransactionDto(ad Domain.TransactionDto) TransactionResponse {
	return TransactionResponse{
		Id: ad.ID,
	}
}
