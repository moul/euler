package challenges

func init() {
	Challenges[31] = Challenge{
		Title:    "Coin sums",
		Correct:  true,
		Callback: C031,
		ID:       31,
	}
}

func c031Rec(coins []int, target int, current *[]int, total *int, idx int) int {
	if idx >= len(coins) {
		return 0
	}
	ways := 0
	for i := 0; *total <= target; i++ {
		if *total == target {
			ways++
			break
		}
		ways += c031Rec(coins, target, current, total, idx+1)
		*total += coins[idx]
		(*current)[idx]++
	}
	*total -= (*current)[idx] * coins[idx]
	(*current)[idx] = 0
	return ways
}

func C031() (interface{}, error) {
	coins := []int{
		200,
		100,
		50,
		20,
		10,
		5,
		2,
		1,
	}
	target := 200
	current := make([]int, len(coins))
	total := 0

	ways := c031Rec(coins, target, &current, &total, 0)

	return ways, nil
}
