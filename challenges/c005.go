package challenges

func init() {
	Challenges[5] = Challenge{
		Title:    "Smallest multiple",
		Correct:  true,
		Callback: C005,
		ID:       5,
	}
}

func C005() (interface{}, error) {
	number := 1

	max := 20
	min := 2

	for i := max; i >= min; i-- {
		if number%i != 0 {
			number *= i
		}
	}

	for i := min; i <= max; i++ {
		if i == 1 {
			continue
		}
		for divisible := true; divisible; {
			for j := min; j <= max; j++ {
				if (number/i)%j != 0 {
					divisible = false
					break
				}
			}
			if divisible {
				number /= i
			}
		}
	}

	return number, nil
}
