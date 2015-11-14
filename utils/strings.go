package utils

import "strings"

func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func LongestRecurringCycleInString(str string) int {
	for size := 1; size < len(str)/2; size++ {
		target := strings.Repeat(str[:size], len(str)/size)
		if strings.Index(str, target) == 0 {
			return size
		}
	}
	return 0
}
