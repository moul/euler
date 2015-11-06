package challenges

func init() {
	Challenges[6] = Challenge{
		Title:    "Sum square difference",
		Correct:  true,
		Callback: C006,
		ID:       6,
	}
}

func C006() (interface{}, error) {
	sum := 0
	squaresSum := 0
	for i := 0; i <= 100; i++ {
		sum += i
		squaresSum += (i * i)
	}
	sumSquare := sum * sum

	return sumSquare - squaresSum, nil
}
