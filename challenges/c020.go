package challenges

import (
	"strconv"

	"moul.io/euler/utils"
)

func init() {
	Challenges[20] = Challenge{
		Title:    "Factorial digit sum",
		Correct:  true,
		Callback: C020,
		ID:       20,
	}
}

func C020() (interface{}, error) {
	sum := 0

	factorial := utils.Factorial(100)
	for _, char := range factorial.String() {
		num, _ := strconv.Atoi(string(char))
		sum += num
	}

	return sum, nil
}
