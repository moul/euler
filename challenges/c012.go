package challenges

import (
	"fmt"

	"github.com/moul/euler/utils"
)

func init() {
	Challenges[12] = Challenge{
		Title:    "Highly divisible triangular number",
		Correct:  true,
		Callback: C012,
		ID:       12,
	}
}

func C012() (interface{}, error) {
	number := 0
	for i := 1; ; i++ {
		number += i
		divisorsCount := utils.DivisorsCount(number)
		if divisorsCount > 500 {
			return number, nil
		}
	}
	return nil, fmt.Errorf("no such solution")
}
