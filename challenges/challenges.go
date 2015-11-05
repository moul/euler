package challenges

import "fmt"

type Challenge struct {
	ID       int
	Title    string
	Correct  bool
	Callback func() (interface{}, error)
}

var Challenges = map[int]Challenge{}

func (c *Challenge) URL() string {
	return fmt.Sprintf("https://projecteuler.net/problem=%d", c.ID)
}
