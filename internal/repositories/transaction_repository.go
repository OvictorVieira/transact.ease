package repositories

import (
	"context"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
)

type transactionRepository struct {
	conn Database
}

const (
	InsertNewTransactionQuery             = `INSERT INTO transact_ease.transactions(account_id, operation_type_id, amount, event_date, created_at, updated_at) VALUES (:account_id, :operation_type_id, :amount, :event_date, :created_at, :updated_at)`
	GetByAccountIdAndOperationTypeIdQuery = `SELECT * FROM transact_ease.transactions WHERE "account_id" = $1 AND "operation_type_id" = $2`
)

func NewTransactionRepository(conn Database) Domain.TransactionRepository {
	return &transactionRepository{
		conn: conn,
	}
}

func (t transactionRepository) Create(ctx context.Context, inTransaction *Domain.TransactionDto) (err error) {
	accountRecord := Domain.FromTransactionDto(inTransaction)

	tx, err := t.conn.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(InsertNewTransactionQuery, accountRecord)

	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (t transactionRepository) GetByAccountIdAndOperationTypeId(ctx context.Context, inTransaction *Domain.TransactionDto) (ouTransaction Domain.TransactionDto, err error) {
	accountRecord := Domain.FromTransactionDto(inTransaction)

	err = t.conn.GetContext(ctx, &accountRecord, GetByAccountIdAndOperationTypeIdQuery, accountRecord.AccountId, accountRecord.OperationTypeId)
	if err != nil {
		return Domain.TransactionDto{}, err
	}

	return accountRecord.ToTransactionDto(), nil
}
