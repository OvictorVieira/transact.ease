package requests

import (
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
)

type AccountCreationRequest struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

func (a *AccountCreationRequest) ToAccountDto() *Domain.AccountDto {
	return &Domain.AccountDto{
		DocumentNumber: a.DocumentNumber,
	}
}

func BuildAccountDtoToFind(accountId string) *Domain.AccountDto {
	return &Domain.AccountDto{
		DocumentNumber: accountId,
	}
}
