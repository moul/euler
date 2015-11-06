package challenges

import "math"

func init() {
	Challenges[206] = Challenge{
		Title:    "Concealed Square",
		Correct:  false,
		Callback: C206,
		ID:       206,
	}
}

func C206() (interface{}, error) {
	min := int(math.Sqrt(10203040506070809))
	max := int(math.Sqrt(19293949596979899)) + 1

	checks := map[int]int{
		9: 1,
		8: 100,
		7: 10000,
		6: 1000000,
		5: 100000000,
		4: 10000000000,
		3: 1000000000000,
		2: 100000000000000,
		1: 10000000000000000,
	}

	for i := min; i < max; i++ {
		pow := i * i
		matched := true
		for target, divisor := range checks {
			if pow/divisor%10 != target {
				matched = false
				break
			}
		}
		if matched {
			return i * 10, nil
		}
	}

	return nil, nil
}
