package utils

func GetDivisorsMethod1(input int) []int {
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

func GetDivisorsMethod2(input int) []int {
	var retLeftSlice []int
	var retRightSlice []int
	var loopMax = input

	retLeftSlice = append(retLeftSlice, 1)
	if input > 1 {
		retRightSlice = append(retRightSlice, input)
	}

	for i := 2; i < loopMax; i++ {
		if input%i == 0 {
			loopMax = input / i
			retLeftSlice = append(retLeftSlice, i)
			if i != loopMax {
				retRightSlice = append(retRightSlice, loopMax)
			}
		}
	}

	for i := len(retRightSlice) - 1; i >= 0; i-- {
		retLeftSlice = append(retLeftSlice, retRightSlice[i])
	}

	return retLeftSlice
}

var GetDivisors = GetDivisorsMethod2

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
