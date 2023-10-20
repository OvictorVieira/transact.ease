package helpers

import (
	"testing"
)

func TestIsArrayContains(t *testing.T) {
	tests := []struct {
		name   string
		arr    []string
		str    string
		result bool
	}{
		{"Contains", []string{"apple", "banana", "cherry"}, "banana", true},
		{"Does not contain", []string{"apple", "banana", "cherry"}, "grape", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsArrayContains(tt.arr, tt.str); got != tt.result {
				t.Errorf("IsArrayContains() = %v, want %v", got, tt.result)
			}
		})
	}
}
