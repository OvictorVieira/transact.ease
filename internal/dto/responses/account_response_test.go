package responses

import (
	"testing"

	Domain "github.com/OvictorVieira/transact.ease/internal/domains/accounts"
	"github.com/stretchr/testify/assert"
)

func TestToAccountDto(t *testing.T) {
	input := &AccountResponse{
		Id:             1,
		DocumentNumber: "123456789",
	}

	output := input.ToAccountDto()

	expected := Domain.AccountDto{
		ID:             1,
		DocumentNumber: "123456789",
	}

	assert.Equal(t, expected, output)
}

func TestFromAccountDto(t *testing.T) {
	input := Domain.AccountDto{
		ID:             1,
		DocumentNumber: "123456789",
	}

	output := FromAccountDto(input)

	expected := AccountResponse{
		Id:             1,
		DocumentNumber: "123456789",
	}

	assert.Equal(t, expected, output)
}
