package main

import (
	"flag"
	"log"
)

func main() {
	token = flags.Get(host, mustToken())

	// tgClient = telegram.New(token);

	// fetcher = fetcher.New()

	// fetcher = fetcher.New()

	//consumer.Start(fetcher, processor)
}

const (
	host = "api.telegram.org"
)

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"Bot token",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("Token not specified")
	}
	return *token
}
