package challenges

import "fmt"

func init() {
	Challenges[23] = Challenge{
		Title:    "Non-abundant sums",
		Correct:  false,
		Callback: C023,
		ID:       23,
	}
}

func GetDivisors(input int) []int {
	ret := []int{}
	for i := 1; i < input/2; i++ {
		if input%i == 0 {
			ret = append(ret, i, input/i)
		}
	}
	return ret
}

func C023() (interface{}, error) {
	fmt.Println(GetDivisors(28))
	return nil, fmt.Errorf("No such answer")
}
