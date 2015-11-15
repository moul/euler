package utils

import (
	"strconv"
	"strings"
)

func IsPandigital(nbStr string) bool {
	length := len(nbStr)
	if length > 9 {
		return false
	}
	for i := 1; i <= length; i++ {
		if !strings.Contains(nbStr, strconv.Itoa(i)) {
			return false
		}
	}
	return true
}
