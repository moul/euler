package challenges

import (
	"fmt"
	"strconv"
)

func init() {
	Challenges[40] = Challenge{
		Title:    "Champernowne's constant",
		Correct:  true,
		Callback: C040,
		ID:       40,
	}
}

func c040_d(bigNumber string, n int) int {
	number := bigNumber[n-1 : n]
	intNumber, _ := strconv.Atoi(number)
	return intNumber
}

func C040() (interface{}, error) {
	bigNumber := ""
	for i := 1; i < 200000; i++ { // FIXME: stop early
		if i%10000 == 0 {
			fmt.Println(len(bigNumber))
		}
		bigNumber += fmt.Sprintf("%d", i)
	}
	fmt.Println(len(bigNumber))

	product := 1
	product *= c040_d(bigNumber, 1)
	product *= c040_d(bigNumber, 10)
	product *= c040_d(bigNumber, 100)
	product *= c040_d(bigNumber, 1000)
	product *= c040_d(bigNumber, 10000)
	product *= c040_d(bigNumber, 100000)
	product *= c040_d(bigNumber, 1000000)

	return product, nil
}
