package main

import (
	"GoTgBot/clients/telegram"
	event_comsumer "GoTgBot/consumer/event-comsumer"
	telegramEvents "GoTgBot/events/telegram"
	"GoTgBot/storage/files"
	"flag"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files-storage"
	batchSize   = 100
)

func main() {
	//token = flags.Get(host, mustToken())

	tgClient := telegram.New(tgBotHost, mustToken())
	eventsProcessor := telegramEvents.New(tgClient, files.New(storagePath))

	log.Print("Service started")

	if err := event_comsumer.New(eventsProcessor, eventsProcessor, batchSize).Start(); err != nil {
		log.Fatal("Service is stopped", err)
	}
	// fetcher = fetcher.New()

	// fetcher = fetcher.New()

	//consumer.Start(fetcher, processor)
}

const (
	host = "api.telegram.org"
)

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"Bot token",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("Token not specified")
	}
	return *token
}
