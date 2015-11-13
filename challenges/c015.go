package challenges

func init() {
	Challenges[15] = Challenge{
		Title:    "Lattice paths",
		Correct:  true,
		Callback: C015,
		ID:       15,
	}
}

func C015() (interface{}, error) {
	width := 21
	height := 21
	c015Map := make([][]int, height)
	for y := 0; y < height; y++ {
		c015Map[y] = make([]int, width)
	}

	c015Map[height-1][width-1] = 1

	for y := height - 1; y >= 0; y-- {
		for x := width - 1; x >= 0; x-- {
			if x > 0 {
				c015Map[y][x-1] += c015Map[y][x]
			}
			if y > 0 {
				c015Map[y-1][x] += c015Map[y][x]
			}
		}
	}
	return c015Map[0][0], nil
}
