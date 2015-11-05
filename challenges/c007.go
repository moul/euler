package challenges

import "github.com/moul/euler/utils"

func init() {
	Challenges[7] = Challenge{
		Title:    "10001st prime",
		Correct:  true,
		Callback: c007,
		ID:       7,
	}
}

func c007() (interface{}, error) {
	// FIXME: use another method
	primes := utils.PrimesLowerThan(1000000)
	return primes[10000], nil
}
