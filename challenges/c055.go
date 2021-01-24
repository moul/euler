package challenges

import (
	"fmt"

	"moul.io/euler/utils"
)

func init() {
	Challenges[55] = Challenge{
		Title:    "Lychrel Numbers",
		Correct:  false,
		Callback: C055,
		ID:       55,
	}
}

func isLychrel(input int) bool {
	last := input
	for i := 0; i < 50; i++ {
		rev := utils.ReverseNumber(last)
		sum := last + rev
		check := utils.IsPalindromeNumber(sum)
		fmt.Println(input, i, last, rev, sum, check)
		if check {
			fmt.Println()
			return true
		}
		last = sum
	}
	return false
}

func C055() (interface{}, error) {
	sum := 0
	for i := 194; i < 10000; i++ {
		check := isLychrel(i)
		fmt.Println(i, check, sum)
		if check {
			sum++
		}
	}
	return sum, nil
}
