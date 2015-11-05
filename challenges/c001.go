package challenges

func init() {
	Challenges[1] = Challenge{
		Title:    "Multiples of 3 and 5",
		Correct:  true,
		Callback: C001,
		ID:       1,
	}
}

func C001() (interface{}, error) {
	sum := 0
	for i := 3; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum, nil
}
