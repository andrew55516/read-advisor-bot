package main

import (
	"flag"
	"log"
	tgClient "read-advisor-bot/clients/telegram"
	event_consumer "read-advisor-bot/consumer/event-consumer"
	"read-advisor-bot/events/telegram"
	"read-advisor-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.NewClient(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Printf("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("empty token")
	}
	return *token
}
