package accounts

import "time"

type AccountDto struct {
	ID             int
	DocumentNumber string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (a *Account) ToAccountDto() AccountDto {
	return AccountDto{
		ID:             a.AccountId,
		DocumentNumber: a.DocumentNumber,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
	}
}

func FromAccountDto(a *AccountDto) Account {
	return Account{
		AccountId:      a.ID,
		DocumentNumber: a.DocumentNumber,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
	}
}
