package challenges

import (
	"math/big"
	"strconv"

	"moul.io/euler/utils"
)

func init() {
	Challenges[34] = Challenge{
		Title:    "Digit factorials",
		Correct:  true,
		Callback: C034,
		ID:       34,
	}
}

func C034() (interface{}, error) {
	total := 0
	for nb := 3; nb < 100000; nb++ {
		factSum := big.NewInt(0)

		for _, digit := range utils.NumberDigits(nb) {
			factSum = factSum.Add(factSum, utils.Factorial(digit))
		}
		if strconv.Itoa(nb) == factSum.String() {
			total += nb
		}
	}
	return total, nil
}
