package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[10] = Challenge{
		Title:    "Summation of primes",
		Correct:  true,
		Callback: C010,
		ID:       10,
	}
}

func C010() (interface{}, error) {
	sum := 0
	for _, prime := range utils.PrimesLowerThan(2000000) {
		sum += prime
	}
	return sum, nil
}
