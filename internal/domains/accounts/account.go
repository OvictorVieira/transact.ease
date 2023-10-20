package accounts

import (
	"time"
)

type Account struct {
	AccountId      int       `db:"account_id"`
	DocumentNumber string    `db:"document_number"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
