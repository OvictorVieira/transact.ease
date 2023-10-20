package requests

import (
	"testing"

	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	"github.com/stretchr/testify/assert"
)

func TestToAccountDto(t *testing.T) {
	input := &AccountCreationRequest{
		DocumentNumber: "123456789",
	}

	output := input.ToAccountDto()

	expected := &Domain.AccountDto{
		DocumentNumber: "123456789",
	}

	assert.Equal(t, expected, output)
}
