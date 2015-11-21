package challenges

import (
	"encoding/csv"
	"os"
	"sort"
)

func init() {
	Challenges[22] = Challenge{
		Title:    "Names scores",
		Correct:  true,
		Callback: C022,
		ID:       22,
	}
}

func AlphabeticalValue(input string) int {
	sum := 0
	for _, rune := range input {
		sum += int(rune) - 64
	}
	return sum
}

func C022() (interface{}, error) {
	input, err := os.Open("./assets/c022_names.txt")
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(input)

	names, err := r.Read()
	if err != nil {
		return nil, err
	}

	sort.Strings(names)

	total := 0
	for i, name := range names {
		total += (i + 1) * AlphabeticalValue(name)
	}

	return total, nil
}
