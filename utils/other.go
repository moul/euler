package utils

func NumberSpiral(size int) [][]int {
	spiral := make([][]int, size)
	for i := 0; i < size; i++ {
		spiral[i] = make([]int, size)
	}

	x := (size - 1) / 2
	y := x
	a, b := 0, -1

	for nb := 1; nb < size*size+1; nb++ {
		spiral[y][x] = nb
		if b == -1 && spiral[y][x+1] == 0 {
			a = 1
			b = 0
		} else if a == 1 && spiral[y+1][x] == 0 {
			a = 0
			b = 1
		} else if b == 1 && spiral[y][x-1] == 0 {
			a = -1
			b = 0
		} else if a == -1 && spiral[y-1][x] == 0 {
			a = 0
			b = -1
		}
		x += a
		y += b
	}

	return spiral
}
