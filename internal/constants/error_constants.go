package constants

import "errors"

var (
	ErrWhenGetFiles  = errors.New("error when get files name")
	ErrWhenReadFiles = errors.New("error when read files name")

	ErrLoadConfig  = errors.New("failed to load config file")
	ErrParseConfig = errors.New("failed to parse env to config struct")
	ErrEmptyVar    = errors.New("required variable environment is empty")

	ErrAccountNotFound          = errors.New("account not found")
	ErrCreateAccount            = errors.New("error when try to create account")
	ErrDuplicatedDocumentNumber = errors.New("error when trying to create account with document number provided")

	ErrTransactionNotFound = errors.New("transaction not found")
	ErrCreatetransaction   = errors.New("error when try to create a transaction")
)
