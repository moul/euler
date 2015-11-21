package utils

func SumIntSlice(slice []int) int {
	sum := 0
	for _, i := range slice {
		sum += i
	}
	return sum
}
