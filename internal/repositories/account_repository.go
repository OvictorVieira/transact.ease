package repositories

import (
	"context"
	"github.com/OvictorVieira/transact.ease/internal/domains"
	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
)

type accountRepository struct {
	conn domains.Database
}

const (
	InsertNewAccountQuery    = `INSERT INTO transact_ease.accounts(document_number, created_at, updated_at) VALUES (:document_number, :created_at, :updated_at)`
	GetByDocumentNumberQuery = `SELECT * FROM transact_ease.accounts WHERE "document_number" = $1`
	GetByAccountIdQuery      = `SELECT * FROM transact_ease.accounts WHERE "account_id" = $1`
)

func NewAccountRepository(conn domains.Database) Domain.AccountRepository {
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

	err = a.conn.GetContext(ctx, &accountRecord, GetByDocumentNumberQuery, accountRecord.DocumentNumber)
	if err != nil {
		return Domain.AccountDto{}, err
	}

	return accountRecord.ToAccountDto(), nil
}

func (a accountRepository) GetById(ctx context.Context, inAccount *Domain.AccountDto) (outAccount Domain.AccountDto, err error) {
	accountRecord := Domain.FromAccountDto(inAccount)

	err = a.conn.GetContext(ctx, &accountRecord, GetByAccountIdQuery, accountRecord.AccountId)
	if err != nil {
		return Domain.AccountDto{}, err
	}

	return accountRecord.ToAccountDto(), nil
}
