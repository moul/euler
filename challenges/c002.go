package challenges

func init() {
	Challenges[2] = Challenge{
		Title:    "Even Fibonacci numbers",
		Correct:  true,
		Callback: C002,
		ID:       2,
	}
}

func C002() (interface{}, error) {
	sum := 0

	left := 1
	right := 2
	for left < 4000000 {
		if right%2 == 0 {
			sum += right
		}
		next := right + left
		left = right
		right = next
	}

	return sum, nil
}
