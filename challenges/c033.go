package challenges

import (
	"math/big"
	"strconv"
	"strings"
)

func init() {
	Challenges[33] = Challenge{
		Title:    "Digit cancelling fractions",
		Correct:  true,
		Callback: C033,
		ID:       33,
	}
}

func LeastCommonMultiple(a, b big.Int) big.Int {
	var z big.Int
	z.Mul(z.Div(&a, z.GCD(nil, nil, &a, &b)), &b)
	return z
}

func C033() (interface{}, error) {
	totalA := int64(1)
	totalB := int64(1)
	for a := 10; a < 100; a++ {
		aStr := strconv.Itoa(a)
		for b := a + 1; b < 100; b++ {
			bStr := strconv.Itoa(b)
			for n := 1; n < 10; n++ {
				nStr := strconv.Itoa(n)
				a2 := strings.Replace(aStr, nStr, "", -1)
				b2 := strings.Replace(bStr, nStr, "", -1)
				if len(a2) != 1 || len(b2) != 1 || a2 == "0" || b2 == "0" {
					continue
				}
				a2nb, _ := strconv.Atoi(a2)
				b2nb, _ := strconv.Atoi(b2)
				if float64(a)/float64(b) == float64(a2nb)/float64(b2nb) {
					totalA *= int64(a)
					totalB *= int64(b)
					// fmt.Println(aStr, bStr, nStr, a2nb, b2nb, float64(a)/float64(b), float64(a2nb)/float64(b2nb))
				}
			}
		}
	}

	rat := big.NewRat(totalA, totalB)

	return rat.Denom(), nil
}
