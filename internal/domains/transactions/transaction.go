package transactions

import (
	"time"
)

type Transaction struct {
	TransactionId   int       `db:"transaction_id"`
	AccountId       int       `db:"account_id"`
	OperationTypeId int       `db:"operation_type_id"`
	Amount          float32   `db:"operation_type_id"`
	EventDate       time.Time `db:"event_date"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
