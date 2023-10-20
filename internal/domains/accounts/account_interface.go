package accounts

import (
	"context"
)

type AccountUsecase interface {
	Create(ctx context.Context, inAccount *AccountDto) (outAccount AccountDto, statusCode int, err error)
	GetById(ctx context.Context, inAccount *AccountDto) (outAccount AccountDto, statusCode int, err error)
}

type AccountRepository interface {
	Create(ctx context.Context, inAccount *AccountDto) (err error)
	GetByDocumentNumber(ctx context.Context, inAccount *AccountDto) (outAccount AccountDto, err error)
	GetById(ctx context.Context, inAccount *AccountDto) (outAccount AccountDto, err error)
}
