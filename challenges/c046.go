package challenges

import (
	"fmt"

	"moul.io/euler/utils"
)

func init() {
	Challenges[46] = Challenge{
		Title:    "Goldbach's other conjecture",
		Correct:  false,
		Callback: C046,
		ID:       46,
	}
}

func C046() (interface{}, error) {
	upperLimit := 10000

	// prepare maps
	primes := utils.PrimesLowerThan(upperLimit)
	primesMap := map[int]bool{}
	twiceSquareMap := []int{}
	for i := 1; ; i++ {
		twiceSquare := 2 * i * i
		if twiceSquare > upperLimit {
			break
		}
		twiceSquareMap = append(twiceSquareMap, twiceSquare)
	}
	for _, prime := range primes {
		primesMap[prime] = true
	}

	// bruteforce
	for i := 3; i < upperLimit; i++ {
		if i%2 == 0 {
			// skipping even numbers
			continue
		}
		if primesMap[i] {
			// skipping prime numbers
			continue
		}

		found := false
		for _, prime := range primes {
			if found {
				break
			}
			if prime > i {
				break
			}
			for _, twiceSquare := range twiceSquareMap {
				sum := prime + twiceSquare
				// fmt.Printf("%d  <->   %d+%d=%d\n", i, prime, twiceSquare, sum)
				if sum == i {
					found = true
					break
				}
				if sum > i {
					break
				}
			}
		}
		if !found {
			return i, nil
		}
	}
	return nil, fmt.Errorf("no such odd composide number bellow %d", upperLimit)
}
