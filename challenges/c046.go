package challenges

func init() {
	Challenges[46] = Challenge{
		Title:    "Goldbach's other conjecture",
		Correct:  false,
		Callback: C046,
		ID:       46,
	}
}

func C046() (interface{}, error) {
	return nil, nil
}
