package challenges

import (
	"strings"

	"github.com/moul/euler/utils"
)

func init() {
	Challenges[24] = Challenge{
		Title:    "Lexicographic permutations",
		Correct:  false,
		Callback: C024,
		ID:       24,
	}
}

func C024() (interface{}, error) {
	input := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	outputs := utils.StringSlicePermutations(input, -1)
	return strings.Join(outputs[999999], ""), nil
}
