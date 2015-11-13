package challenges

import (
	"fmt"

	"github.com/moul/euler/utils"
)

func init() {
	Challenges[23] = Challenge{
		Title:    "Non-abundant sums",
		Correct:  false,
		Callback: C023,
		ID:       23,
	}
}

func C023() (interface{}, error) {
	fmt.Println(utils.GetDivisors(28))
	return nil, fmt.Errorf("No such answer")
}
