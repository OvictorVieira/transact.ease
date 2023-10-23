package responses

import (
	"testing"

	Domain "github.com/OvictorVieira/transact.ease/internal/domains/transactions"
	"github.com/stretchr/testify/assert"
)

func TestToTransactionDto(t *testing.T) {
	input := &TransactionResponse{
		Id: 1,
	}

	output := input.ToTransactionDto()

	expected := Domain.TransactionDto{
		ID: 1,
	}

	assert.Equal(t, expected, output)
}

func TestFromTransactionDto(t *testing.T) {
	input := Domain.TransactionDto{
		ID: 1,
	}

	output := FromTransactionDto(input)

	expected := TransactionResponse{
		Id: 1,
	}

	assert.Equal(t, expected, output)
}
