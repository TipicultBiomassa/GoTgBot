package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()
	// token = flags.Get(token)

	// tgClient = telegram.New(token);

	// fetcher = fetcher.New()

	// fetcher = fetcher.New()

	//consumer.Start(fetcher, processor)
}

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
