package challenges

import "github.com/moul/euler/utils"

func init() {
	Challenges[3] = Challenge{
		Title:    "Largest prime factor",
		Correct:  true,
		Callback: c003,
		ID:       3,
	}
}

func c003() (interface{}, error) {
	factors := utils.PrimeFactors(600851475143)
	return factors[len(factors)-1], nil
}
