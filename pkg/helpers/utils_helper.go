package helpers

import (
	"strconv"
)

func IsArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
