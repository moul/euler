package challenges

import (
	"fmt"
	"math/big"

	"github.com/moul/euler/utils"
)

func init() {
	Challenges[26] = Challenge{
		Title:    "Reciprocal cycles",
		Correct:  false,
		Callback: C026,
		ID:       26,
	}
}

func C026() (interface{}, error) {
	longestSize := 0
	var longestNumber int64

	for i := int64(2); i < 1000; i++ {
		rat := big.NewRat(1, i)
		floatString := rat.FloatString(500)
		cycleSize := utils.LongestRecurringCycleInString(utils.ReverseString(floatString[10:]))

		// fmt.Println(i, floatString, cycleSize)
		if cycleSize > longestSize {
			fmt.Println(i, cycleSize, floatString)
			longestSize = cycleSize
			longestNumber = i
		}
	}

	return longestNumber, nil
}
