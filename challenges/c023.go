package challenges

import "fmt"

func init() {
	Challenges[23] = Challenge{
		Title:    "Non-abundant sums",
		Correct:  false,
		Callback: C023,
		ID:       23,
	}
}

func C023() (interface{}, error) {
	return nil, fmt.Errorf("No such answer")
}
