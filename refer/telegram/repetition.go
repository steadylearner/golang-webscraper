package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	err := godotenv.Load()

	// What is difference when used with return or without it?
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	telegramApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	// fmt.Println(telegramApiToken)

	b, err := tb.NewBot(tb.Settings{
		Token:  telegramApiToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	// 1. Make repetition.(Done) https://godoc.org/gopkg.in/tucnak/telebot.v2#Message
	// 2. Parse as html. https://github.com/tucnak/telebot#send-options
	b.Handle(tb.OnText, func(m *tb.Message) {
		fmt.Println(m.Text)
		b.Send(m.Sender, m.Text)
	})

	// b.Handle("/hello", func(m *tb.Message) {
	// 	b.Send(m.Sender, "hello world")
	// })

	b.Start()
}
