package challenges

import (
	"strconv"

	"moul.io/euler/utils"
)

func init() {
	Challenges[41] = Challenge{
		Title:    "Pandigital prime",
		Correct:  true,
		Callback: C041,
		ID:       41,
	}
}

func C041() (interface{}, error) {
	max := 0
	upperLimit := 987654321
	primes := utils.PrimesLowerThan(upperLimit)
	for _, prime := range primes {
		if utils.IsPandigital(strconv.Itoa(prime)) {
			if prime > max {
				max = prime
				// fmt.Println(prime)
			}
		}
	}
	return max, nil
}
