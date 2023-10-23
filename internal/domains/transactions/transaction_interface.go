package transactions

import (
	"context"
)

type TransactionUsecase interface {
	Create(ctx context.Context, inTransaction *TransactionDto) (ouTransaction TransactionDto, statusCode int, err error)
}

type TransactionRepository interface {
	Create(ctx context.Context, inTransaction *TransactionDto) (err error)
	GetByAccountIdAndOperationTypeId(ctx context.Context, inTransaction *TransactionDto) (ouTransaction TransactionDto, err error)
}
