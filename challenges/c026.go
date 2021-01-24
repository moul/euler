package challenges

import (
	"math/big"

	"moul.io/euler/utils"
)

func init() {
	Challenges[26] = Challenge{
		Title:    "Reciprocal cycles",
		Correct:  true,
		Callback: C026,
		ID:       26,
	}
}

// FIXME: use loop which divide and modulo instead of parsing strings

func C026() (interface{}, error) {
	longestSize := 1
	var longestNumber int64

	for i := int64(2); i < 1000; i++ {
		rat := big.NewRat(1, i)
		floatString := rat.FloatString(2000)
		cycleSize := utils.LongestRecurringCycleInString(utils.ReverseString(floatString[10:]), 1)

		// fmt.Println(i, floatString, cycleSize)
		if cycleSize > longestSize {
			// fmt.Println(i, cycleSize, floatString)
			longestSize = cycleSize
			longestNumber = i
		}
	}

	return longestNumber, nil
}
