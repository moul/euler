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

func IntPow(a, b int) int {
	nb := 1
	for i := 0; i < b; i++ {
		nb *= a
	}
	return nb
}

func NumberDigits(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}
