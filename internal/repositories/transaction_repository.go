package repositories

import (
	"github.com/OvictorVieira/transact.ease/internal/domains"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
)

type transactionRepository struct {
	conn domains.Database
}

const (
	InsertNewTransactionQuery = `INSERT INTO transact_ease.transactions(account_id, operation_type_id, amount, event_date, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING transaction_id`
	NonExistentId             = -1
)

func NewTransactionRepository(conn domains.Database) Domain.TransactionRepository {
	return &transactionRepository{
		conn: conn,
	}
}

func (t transactionRepository) Create(inTransaction *Domain.TransactionDto) (lastInsertId int, err error) {
	transactionRecord := Domain.FromTransactionDto(inTransaction)

	tx, err := t.conn.Begin()
	if err != nil {
		return NonExistentId, err
	}

	err = tx.QueryRow(
		InsertNewTransactionQuery,
		transactionRecord.AccountId,
		transactionRecord.OperationTypeId,
		transactionRecord.Amount,
		transactionRecord.EventDate,
		transactionRecord.CreatedAt,
		transactionRecord.UpdatedAt,
	).Scan(&lastInsertId)

	if err != nil {
		_ = tx.Rollback()
		return NonExistentId, err
	}

	err = tx.Commit()
	if err != nil {
		return NonExistentId, err
	}

	return lastInsertId, nil
}
