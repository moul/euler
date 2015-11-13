package utils

func CollatzLengthMethod1(number int) int {
	count := 1
	for number > 1 {
		count++
		if number%2 == 0 {
			number /= 2
		} else {
			number = number*3 + 1
		}
	}
	return count
}

var CollatzLengthCache map[int]int

func init() {
	CollatzLengthCache = make(map[int]int, 0)
}

func CollatzLengthMethod2(base int) int {
	number := base
	count := 1
	for number > 1 {
		count++
		if number%2 == 0 {
			number /= 2
		} else {
			number = number*3 + 1
		}
		if val, found := CollatzLengthCache[number]; found {
			count += val
			break
		}
	}
	CollatzLengthCache[base] = count
	return count
}

func CollatzLengthMethod3(number int) int {
	if number == 1 {
		return 1
	}

	if val, found := CollatzLengthCache[number]; found {
		return val
	}

	var count int
	if number%2 == 0 {
		count = CollatzLengthMethod3(number/2) + 1
	} else {
		count = CollatzLengthMethod3(number*3+1) + 1
	}
	CollatzLengthCache[number] = count
	return count
}

var CollatzLength = CollatzLengthMethod1
