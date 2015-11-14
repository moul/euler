package challenges

func init() {
	Challenges[39] = Challenge{
		Title:    "Integer right triangles",
		Correct:  true,
		Callback: C039,
		ID:       39,
	}
}

func C039() (interface{}, error) {
	solutionsMax := 0
	pMax := 0

	for p := 1; p < 1000; p++ {
		solutions := 0
		for a := 1; a < p; a++ {
			for b := a; b < p-a; b++ {
				c := p - b - a
				a2 := a * a
				b2 := b * b
				c2 := c * c
				if a2+b2 == c2 {
					solutions++
				}
			}
		}
		if solutions > solutionsMax {
			solutionsMax = solutions
			pMax = p
		}
	}

	return pMax, nil
}
