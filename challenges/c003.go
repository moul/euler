package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[3] = Challenge{
		Title:    "Largest prime factor",
		Correct:  true,
		Callback: C003,
		ID:       3,
	}
}

func C003() (interface{}, error) {
	factors := utils.PrimeFactors(600851475143)
	return factors[len(factors)-1], nil
}
