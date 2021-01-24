package challenges

import (
	"fmt"
	"math/big"

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

func isLychrel(input int64) bool {
	last := big.NewInt(input)
	for i := 0; i < 50; i++ {
		rev := utils.ReverseBigInt(last)
		sum := big.NewInt(0)
		sum.Add(last, rev)
		check := utils.IsPalindromeBigInt(sum)
		// fmt.Println(input, i, last, rev, sum, check)
		if check {
			return false
		}
		last = sum
	}
	return true
}

func C055() (interface{}, error) {
	sum := 0
	for i := 0; i < 10000; i++ {
		check := isLychrel(int64(i))
		fmt.Println(i, check, sum)
		if check {
			sum++
		}
	}
	return sum, nil
}
