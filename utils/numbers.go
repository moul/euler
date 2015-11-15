package utils

import (
	"strconv"
	"strings"
)

func IsPandigital(nbStr string) bool {
	if len(nbStr) != 9 {
		return false
	}
	for i := 1; i <= 9; i++ {
		if !strings.Contains(nbStr, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}
