package utils

import "math"

func PrimesLowerThan(n int) []int {
	if n == 2 {
		return []int{2}
	} else if n < 2 {
		return []int{}
	}

	s := []int{}
	for i := 3; i < n+1; i += 2 {
		s = append(s, i)
	}

	root := math.Sqrt(float64(n))
	half := (n+1)/2 - 1
	i := 0
	m := 3

	for float64(m) <= root {
		if s[i] != 0 {
			j := (m*m - 3) / 2
			s[j] = 0
			for j < half {
				s[j] = 0
				j += m
			}
		}
		i++
		m = 2*i + 3
	}

	output := []int{2}
	for i := 0; i < len(s); i++ {
		if s[i] > 0 {
			output = append(output, s[i])
		}
	}
	return output
}
