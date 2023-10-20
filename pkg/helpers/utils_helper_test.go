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

func TestStringToFloat(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{"Valid float", "123.45", 123.45, false},
		{"Invalid float", "abc", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToFloat(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
