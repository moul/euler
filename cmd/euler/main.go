package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/jessevdk/go-flags"

	"moul.io/euler/challenges"
)

func main() {
	var opts struct {
		ChallengeID int `short:"c" long:"challenge" description:"Challenge ID" required:"true"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		logrus.Fatalf("Failed to parse arguments: %v", err)
	}

	challenge, found := challenges.Challenges[opts.ChallengeID]
	if !found {
		logrus.Fatalf("No such challenge %d", opts.ChallengeID)
	}

	fmt.Println(challenge.Title)
	fmt.Println(strings.Repeat("=", len(challenge.Title)))
	fmt.Println()

	fmt.Printf("URL: %s\n", challenge.URL())
	t0 := time.Now()
	ret, err := challenge.Callback()
	t1 := time.Now()
	fmt.Printf("Duration: %v\n", t1.Sub(t0))

	if err != nil {
		logrus.Fatalf("Failed to execute challenge: %v", err)
	}
	fmt.Printf("Answer: %v\n", ret)
}
