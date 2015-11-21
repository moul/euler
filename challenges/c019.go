package challenges

import "time"

func init() {
	Challenges[19] = Challenge{
		Title:    "Counting Sundays",
		Correct:  true,
		Callback: C019,
		ID:       19,
	}
}

func C019() (interface{}, error) {
	sundays := 0

	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		return nil, err
	}

	for year := 1901; year <= 2000; year++ {
		for month := time.January; month <= time.December; month++ {
			date := time.Date(year, month, 1, 0, 0, 0, 0, loc)
			if date.Weekday() == 6 {
				sundays++
			}
		}
	}

	return sundays, nil
}
