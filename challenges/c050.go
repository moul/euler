package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[50] = Challenge{
		Title:    "Consecutive prime sum",
		Correct:  true,
		Callback: C050,
		ID:       50,
	}
}

func C050() (interface{}, error) {
	limit := 1000000
	primes := utils.PrimesLowerThan(limit)
	totalLength := len(primes)
	asKeys := map[int]bool{}
	for _, prime := range primes {
		asKeys[prime] = true
	}

	max := 0
	maxValue := 0
	for start := 0; start < totalLength; start++ {
		sum := 0
		for length := 0; length < totalLength-start; length++ {
			sum += primes[start+length]
			if _, found := asKeys[sum]; found {
				if max < length+1 {
					// fmt.Println(primes[start:start+length+1], length+1, sum)
					max = length + 1
					maxValue = sum
				}
			}
			if sum >= limit {
				break
			}
		}
	}
	return maxValue, nil
}
