package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[23] = Challenge{
		Title:    "Non-abundant sums",
		Correct:  true,
		Callback: C023,
		ID:       23,
	}
}

func C023() (interface{}, error) {
	upperLimit := 28123

	numbers := map[int]bool{}
	for i := 1; i < upperLimit; i++ {
		numbers[i] = true
	}

	abundantNumbers := []int{}
	for i := 2; i < upperLimit; i++ {
		divisors := utils.GetDivisors(i)
		sumDivisors := 0
		for _, divisor := range divisors[:len(divisors)-1] {
			sumDivisors += divisor
		}
		if sumDivisors > i {
			abundantNumbers = append(abundantNumbers, i)
		}
	}

	abundantNumbersLen := len(abundantNumbers)

	for i := 0; i < abundantNumbersLen; i++ {
		for j := i; j < abundantNumbersLen; j++ {
			abundantNumbersSum := abundantNumbers[i] + abundantNumbers[j]
			if _, found := numbers[abundantNumbersSum]; found {
				numbers[abundantNumbersSum] = false
			}
		}
	}

	sumNonAbundant := 0
	for number, isNotSum := range numbers {
		if isNotSum {
			sumNonAbundant += number
		}
	}

	return sumNonAbundant, nil
}
