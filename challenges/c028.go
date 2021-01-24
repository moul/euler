package challenges

import "moul.io/euler/utils"

func init() {
	Challenges[28] = Challenge{
		Title:    "Number spiral diagonals",
		Correct:  true,
		Callback: C028,
		ID:       28,
	}
}

func C028() (interface{}, error) {
	size := 1001
	spiral := utils.NumberSpiral(size)
	sum := -1
	for i := 0; i < size; i++ {
		sum += spiral[i][i]
		sum += spiral[i][size-i-1]
	}
	return sum, nil
}
