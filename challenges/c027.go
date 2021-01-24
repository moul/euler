package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[27] = Challenge{
		Title:    "Quadratic primes",
		Correct:  true,
		Callback: C027,
		ID:       27,
	}
}

func C027() (interface{}, error) {
	maxA, maxB, maxN := 0, 0, 0
	for a := -1000; a < 1000; a++ {
		for b := 0; b < 1000; b++ {
			n := 0
			for utils.IsPrime(n*n + a*n + b) {
				n++
			}
			if n > maxN {
				maxN = n
				maxA = a
				maxB = b
			}
		}
	}
	return maxA * maxB, nil
}
