package challenges

import "math/big"

func init() {
	Challenges[16] = Challenge{
		Title:    "Power digit sum",
		Correct:  true,
		Callback: C016,
		ID:       16,
	}
}

func C016() (interface{}, error) {
	two := big.NewInt(2)
	nb := big.NewInt(1)
	for i := 0; i < 1000; i++ {
		nb.Mul(nb, two)
	}

	sum := 0

	for _, figure := range nb.String() {
		sum += int(figure - 48)
	}

	return sum, nil
}
