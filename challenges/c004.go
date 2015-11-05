package challenges

import (
	"strconv"

	"github.com/moul/euler/utils"
)

func init() {
	Challenges[4] = Challenge{
		Title:    "Largest palindrome product",
		Correct:  true,
		Callback: C004,
		ID:       4,
	}
}

func C004() (interface{}, error) {
	biggest := 0

	for a := 999; a >= 100; a-- {
		for b := 999; b >= 100; b-- {
			product := a * b
			if product < biggest {
				continue
			}
			productStr := strconv.Itoa(product)
			if productStr == utils.ReverseString(productStr) {
				if biggest < product {
					biggest = product
				}
			}
		}
	}

	return biggest, nil
}
