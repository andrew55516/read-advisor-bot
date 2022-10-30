package main

import (
	"flag"
	"log"
	"read-advisor-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.NewClient(tgBotHost, mustToken())

	//fetcher := fetcher.New(tgClient)

	//processor := processor.New(tgClient)

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("empty token")
	}
	return *token
}
