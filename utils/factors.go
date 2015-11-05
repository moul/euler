package utils

import "math/big"

func PrimeFactors(input int) []int {
	output := []int{}

	current := input

	for i := 2; i < current; i++ {
		if current%i == 0 {
			current /= i
			output = append(output, i)
		}
	}
	output = append(output, current)

	return output
}

func Factorial(input int) *big.Int {
	x := new(big.Int)
	x.MulRange(1, int64(input))
	return x
}
