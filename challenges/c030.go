package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[30] = Challenge{
		Title:    "Digit fifth powers",
		Correct:  true,
		Callback: C030,
		ID:       30,
	}
}

func C030() (interface{}, error) {
	total := 0
	for nb := 2; nb < 1000000; nb++ {
		sum := 0
		for _, digit := range utils.NumberDigits(nb) {
			sum += utils.IntPow(digit, 5)
		}
		if sum == nb {
			total += sum
		}
	}
	return total, nil
}
