package main

import (
	"log"
	"os"
	"strconv"

	"github.com/broady/go-travis"
)

func main() {
	tok := os.Getenv("TRAVIS_API_TOKEN")
	if tok == "" {
		log.Fatal("must have TRAVIS_API_TOKEN set")
	}

	c := travis.NewDefaultClient(tok)

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	build, _, commit, _, err := c.Builds.Get(uint(id))
	if err != nil {
		log.Fatal(err)
	}
	log.Print(build.EventType)
	log.Print(commit.Message)
}
