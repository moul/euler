package challenges

import (
	"fmt"

	"moul.io/euler/utils"
)

func init() {
	Challenges[32] = Challenge{
		Title:    "Pandigital products",
		Correct:  true,
		Callback: C032,
		ID:       32,
	}
}

func C032() (interface{}, error) {
	products := map[int]bool{}
	for a := 1; a < 9000; a++ {
		for b := a + 1; b < 9000; b++ {
			product := a * b
			nbStr := fmt.Sprintf("%d%d%d", a, b, product)
			length := len(nbStr)
			if length > 9 {
				break
			}
			if length != 9 {
				continue
			}
			if utils.IsPandigital(nbStr) {
				products[product] = true
			}
		}
	}
	sum := 0
	for product := range products {
		sum += product
	}
	return sum, nil
}
