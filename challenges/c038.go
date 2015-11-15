package challenges

import (
	"strconv"

	"github.com/moul/euler/utils"
)

func init() {
	Challenges[38] = Challenge{
		Title:    "Pandigital multiples",
		Correct:  true,
		Callback: C038,
		ID:       38,
	}
}

func C038() (interface{}, error) {
	max := 0
	for left := 1; left < 10000; left++ {
		output := ""
		for right := 1; right < 9 && len(output) < 9; right++ {
			output += strconv.Itoa(left * right)
		}
		if utils.IsPandigital(output) {
			nb, _ := strconv.Atoi(output)
			// fmt.Println(left, nb)
			if nb > max {
				max = nb
			}
		}
	}
	return max, nil
}
