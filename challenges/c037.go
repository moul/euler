package challenges

import (
	"strconv"

	"github.com/moul/euler/utils"
)

func init() {
	Challenges[37] = Challenge{
		Title:    "Truncatable primes",
		Correct:  true,
		Callback: C037,
		ID:       37,
	}
}

func isStrPrime(nbStr string) bool {
	nb, _ := strconv.Atoi(nbStr)
	return utils.IsPrime(nb)
}

func C037() (interface{}, error) {
	sum := 0
	for _, prime := range utils.PrimesLowerThan(1000000) {
		if prime < 10 {
			continue
		}
		primeStr := strconv.Itoa(prime)
		ok := true
		for i := 1; i < len(primeStr); i++ {
			if !isStrPrime(primeStr[i:]) || !isStrPrime(primeStr[:i]) {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		sum += prime
	}
	return sum, nil
}
