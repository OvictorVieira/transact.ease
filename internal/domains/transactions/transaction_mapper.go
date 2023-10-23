package transactions

import "time"

type TransactionDto struct {
	ID              int
	AccountId       int
	OperationTypeId int
	Amount          float32
	EventDate       time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (a *Transaction) ToTransactionDto() TransactionDto {
	return TransactionDto{
		ID:              a.TransactionId,
		AccountId:       a.AccountId,
		OperationTypeId: a.OperationTypeId,
		Amount:          a.Amount,
		EventDate:       a.EventDate,
		CreatedAt:       a.CreatedAt,
		UpdatedAt:       a.UpdatedAt,
	}
}

func FromTransactionDto(a *TransactionDto) Transaction {
	return Transaction{
		TransactionId:   a.ID,
		AccountId:       a.AccountId,
		OperationTypeId: a.OperationTypeId,
		Amount:          a.Amount,
		EventDate:       a.EventDate,
		CreatedAt:       a.CreatedAt,
		UpdatedAt:       a.UpdatedAt,
	}
}
