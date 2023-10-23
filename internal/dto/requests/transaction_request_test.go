package requests

import (
	"testing"

	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
	"github.com/stretchr/testify/assert"
)

func TestToTransactionDto(t *testing.T) {
	input := &TransactionCreationRequest{
		AccountId:       1,
		OperationTypeId: 1,
		Amount:          52,
	}

	output := input.ToTransactionDto()

	expected := &Domain.TransactionDto{
		AccountId:       1,
		OperationTypeId: 1,
		Amount:          52,
	}

	assert.Equal(t, expected, output)
}
