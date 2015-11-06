package challenges

import "fmt"

func init() {
	Challenges[9] = Challenge{
		Title:    "Special Pythagorean triplet",
		Correct:  true,
		Callback: C009,
		ID:       9,
	}
}

func C009() (interface{}, error) {
	target := 1000

	for a := 1; a < target-4; a++ {
		pa := a * a
		for b := a + 1; b < target-3; b++ {
			pb := b * b
			for c := b + 1; c < target-2; c++ {
				if a+b+c != target {
					continue
				}
				pc := c * c
				if pc == pa+pb {
					return a * b * c, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("No such answer")
}
