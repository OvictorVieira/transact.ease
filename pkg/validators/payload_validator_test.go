package validators

import (
	"testing"
)

// MockPayload for testing purposes
type MockPayload struct {
	Field string `json:"field" validate:"required"`
}

func TestValidatePayloads_WithRequiredField(t *testing.T) {
	payload := &MockPayload{Field: "value"}
	err := ValidatePayloads(payload)
	if err != nil {
		t.Errorf("ValidatePayloads() unexpected error = %v", err)
	}
}

func TestValidatePayloads_WithoutRequiredField(t *testing.T) {
	payload := &MockPayload{}
	err := ValidatePayloads(payload)
	if err == nil {
		t.Errorf("ValidatePayloads() error = nil, wantErr true")
		return
	}
	expectedErrMsg := "field: is a required field"
	if err.Error() != expectedErrMsg {
		t.Errorf("ValidatePayloads() error = %v, wantErr %v", err, expectedErrMsg)
	}
}
