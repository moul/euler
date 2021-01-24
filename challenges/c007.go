package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[7] = Challenge{
		Title:    "10001st prime",
		Correct:  true,
		Callback: C007,
		ID:       7,
	}
}

func C007() (interface{}, error) {
	// FIXME: use another method
	primes := utils.PrimesLowerThan(1000000)
	return primes[10000], nil
}
