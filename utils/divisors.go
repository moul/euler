package utils

func GetDivisors(input int) []int {
	// FIXME: remove duplicates
	// FIXME: sort results
	ret := []int{}
	for i := 1; i < input/2; i++ {
		if input%i == 0 {
			ret = append(ret, i, input/i)
		}
	}
	return ret
}

func DivisorsCount(input int) int {
	if input == 1 {
		return 1
	}
	divisor := 0
	number := input
	for i := 1; i < number; i++ {
		if input%i == 0 {
			divisor += 2
			if i*i == input {
				divisor--
			}
			number = input / i
		}
	}
	return divisor
}
