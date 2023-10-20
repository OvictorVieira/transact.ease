package repositories

import (
	"context"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	"github.com/jmoiron/sqlx"
)

type accountRepository struct {
	conn *sqlx.DB
}

const (
	InsertNewAccountQuery = `INSERT INTO transact_ease.accounts(document_number, created_at, updated_at) VALUES (:document_number, :created_at, :updated_at)`
	GetByIdQuery          = `SELECT * FROM transact_ease.accounts WHERE "document_number" = $1`
)

func NewAccountRepository(conn *sqlx.DB) Domain.AccountRepository {
	return &accountRepository{
		conn: conn,
	}
}

func (a accountRepository) Create(ctx context.Context, inAccount *Domain.AccountDto) (err error) {
	accountRecord := Domain.FromAccountDto(inAccount)

	_, err = a.conn.NamedQueryContext(ctx, InsertNewAccountQuery, accountRecord)
	if err != nil {
		return err
	}

	return nil
}

func (a accountRepository) GetByDocumentNumber(ctx context.Context, inAccount *Domain.AccountDto) (outAccount Domain.AccountDto, err error) {
	accountRecord := Domain.FromAccountDto(inAccount)

	err = a.conn.GetContext(ctx, &accountRecord, GetByIdQuery, accountRecord.DocumentNumber)
	if err != nil {
		return Domain.AccountDto{}, err
	}

	return accountRecord.ToAccountDto(), nil
}
