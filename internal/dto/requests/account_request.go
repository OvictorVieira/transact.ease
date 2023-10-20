package requests

import (
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
)

type AccountCreationRequest struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

func (u *AccountCreationRequest) ToAccountDto() *Domain.AccountDto {
	return &Domain.AccountDto{
		DocumentNumber: u.DocumentNumber,
	}
}
