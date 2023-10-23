package transactions

type TransactionUsecase interface {
	Create(inTransaction *TransactionDto) (ouTransaction *TransactionDto, statusCode int, err error)
}

type TransactionRepository interface {
	Create(inTransaction *TransactionDto) (createdTransactionId int, err error)
}
