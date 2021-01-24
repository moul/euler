package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[92] = Challenge{
		Title:    "Square digit chains",
		Correct:  true,
		Callback: C092,
		ID:       92,
	}
}

func c092GetLoopNumber(starting int) int {
	for starting != 1 && starting != 89 {
		sum := 0
		for _, digit := range utils.NumberDigits(starting) {
			sum += digit * digit
		}
		starting = sum
	}
	return starting
}

func C092() (interface{}, error) {
	count89 := 0
	for i := 1; i < 10000000; i++ {
		if c092GetLoopNumber(i) == 89 {
			count89++
		}
	}
	return count89, nil
}
