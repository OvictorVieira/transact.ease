package responses

import (
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
)

type AccountResponse struct {
	Id             int    `json:"id"`
	DocumentNumber string `json:"document_number"`
}

func (ar *AccountResponse) ToAccountDto() Domain.AccountDto {
	return Domain.AccountDto{
		ID:             ar.Id,
		DocumentNumber: ar.DocumentNumber,
	}
}

func FromAccountDto(ad Domain.AccountDto) AccountResponse {
	return AccountResponse{
		Id:             ad.ID,
		DocumentNumber: ad.DocumentNumber,
	}
}
