package challenges

import (
	"strconv"
	"strings"
)

func init() {
	Challenges[18] = Challenge{
		Title:    "Maximum path sum I",
		Correct:  true,
		Callback: C018,
		ID:       18,
	}
}

var c018Str = strings.TrimSpace(`
75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
`)

func C018() (interface{}, error) {
	c018Triangle := make([][]int, 0)
	for _, lineStr := range strings.Split(c018Str, "\n") {
		line := []int{}
		for _, nbStr := range strings.Split(lineStr, " ") {
			nb, _ := strconv.Atoi(nbStr)
			line = append(line, nb)
		}
		c018Triangle = append(c018Triangle, line)
	}

	// fmt.Println(c018Triangle)

	for y := len(c018Triangle) - 2; y > -1; y-- {
		for x := 0; x < len(c018Triangle[y]); x++ {
			max := 0
			if c018Triangle[y+1][x] > max {
				max = c018Triangle[y+1][x]
			}
			if c018Triangle[y+1][x+1] > max {
				max = c018Triangle[y+1][x+1]
			}
			c018Triangle[y][x] += max
		}
	}

	// fmt.Println(c018Triangle)

	return c018Triangle[0][0], nil
}
