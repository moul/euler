package challenges

import (
	"fmt"
	"math/big"
)

func init() {
	Challenges[25] = Challenge{
		Title:    "1000-digit Fibonacci number",
		Correct:  true,
		Callback: C025,
		ID:       25,
	}
}

func C025() (interface{}, error) {
	left := big.NewInt(1)
	right := big.NewInt(2)
	for i := 2; i < 10000; i++ {
		next := big.NewInt(0)
		next.Add(next, left)
		next.Add(next, right)
		if len(left.String()) > 999 {
			return i, nil
		}
		left = right
		right = next
	}
	return nil, fmt.Errorf("No such result")
}
