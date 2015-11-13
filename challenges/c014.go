package challenges

import "github.com/moul/euler/utils"

func init() {
	Challenges[14] = Challenge{
		Title:    "Longest Collatz sequence",
		Correct:  true,
		Callback: C014,
		ID:       14,
	}
}

func C014() (interface{}, error) {
	longestNb := 0
	longestLength := 0
	for i := 1; i < 1000000; i++ {
		length := utils.CollatzLength(i)
		if length > longestLength {
			longestLength = length
			longestNb = i
		}
	}
	return longestNb, nil
}
