package challenges

import (
	"strconv"

	"github.com/moul/euler/utils"
)

func init() {
	Challenges[36] = Challenge{
		Title:    "Double-base palindromes",
		Correct:  true,
		Callback: C036,
		ID:       36,
	}
}

func C036() (interface{}, error) {
	sum := 0
	for nb := 1; nb < 1000000; nb++ {
		if nb%2 == 0 {
			continue
		}
		nbStr := strconv.Itoa(nb)
		if nbStr != utils.ReverseString(nbStr) {
			continue
		}
		binary := strconv.FormatInt(int64(nb), 2)
		if binary != utils.ReverseString(binary) {
			continue
		}
		sum += nb
	}
	return sum, nil
}
