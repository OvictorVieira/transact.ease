package helpers

import (
	"path/filepath"
	"strconv"
	"strings"
)

func ExtractLeadingNumber(filename string) int {
	base := filepath.Base(filename)
	parts := strings.Split(base, "_")
	if len(parts) == 0 {
		return 0
	}
	num, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0
	}
	return num
}
