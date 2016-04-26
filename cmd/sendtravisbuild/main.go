package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/broady/go-travis"
)

var tok = flag.String("tok", "", "travis auth token. required")
var slug = flag.String("slug", "broady/travisplay", "slug. required")

func main() {
	flag.Parse()

	if *tok == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	const body = `
	{
		"request": {
			"type": "cron",
			"event_type": "cron",
			"message": "Build started by ci-aeon."
		}
	}
	`

	req, err := http.NewRequest("POST", travis.TRAVIS_API_DEFAULT_URL, strings.NewReader(body))
	if err != nil {
		log.Fatalf("create req: %v", err)
	}
	req.URL.Opaque = fmt.Sprintf("/repo/%s/requests", url.QueryEscape(*slug))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Travis-API-Version", "3")
	req.Header.Add("Authorization", "token "+*tok)
	req.Header.Add("User-Agent", "TravisCron/0.1 (github.com/broady)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range resp.Header {
		fmt.Printf("%v %v\n", k, v)
	}
	io.Copy(os.Stdout, resp.Body)
}
