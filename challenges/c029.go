package challenges

import "math/big"

func init() {
	Challenges[29] = Challenge{
		Title:    "Distinct powers",
		Correct:  true,
		Callback: C029,
		ID:       29,
	}
}

func C029() (interface{}, error) {
	numbers := map[string]bool{}
	for a := 2; a <= 100; a++ {
		bigA := big.NewInt(int64(a))
		nb := big.NewInt(1)
		nb.Mul(nb, bigA)
		for b := 2; b <= 100; b++ {
			nb.Mul(nb, bigA)
			numbers[nb.String()] = true
		}
	}
	return len(numbers), nil
}
