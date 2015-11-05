package challenges

import (
	"fmt"
	"math/big"
)

func init() {
	Challenges[26] = Challenge{
		Title:    "Reciprocal cycles",
		Correct:  false,
		Callback: c026,
		ID:       26,
	}
}

func c026() (interface{}, error) {
	for i := int64(2); i < int64(1000); i++ {
		rat := big.NewRat(1, i)
		fmt.Println(i, rat.FloatString(500))
	}

	return nil, nil
}
