package main

import (
	"fmt"

	"github.com/t0mmyt/slacker"
)

const SlackURL = "https://hooks.slack.com/services/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"

// const SlackURL = "http://localhost:5000"

func main() {
	m := slacker.Message{
		Username: "maki",
		Icon:     ":shit:",
		Text:     "There's a first time for everything",
	}

	s, err := slacker.NewSlackEndpoint(SlackURL)
	if err != nil {
		panic(err)
	}

	r, err := s.Send(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)

}
