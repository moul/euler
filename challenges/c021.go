package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[21] = Challenge{
		Title:    "Amicable numbers",
		Correct:  true,
		Callback: C021,
		ID:       21,
	}
}

func c021_d(input int) int {
	divisors := utils.GetDivisors(input)
	return utils.SumIntSlice(divisors) - input
}

func C021() (interface{}, error) {
	sum := 0

	upperLimit := 10000

	for a := 1; a < 10000-1; a++ {
		b := c021_d(a)
		if b >= upperLimit || b <= a {
			continue
		}

		if c021_d(b) == a {
			sum += a + b
		}
	}

	return sum, nil
}
