package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/wdipax/match/adapter/tgevent"
	"github.com/wdipax/match/adapter/tgresponse"
	"github.com/wdipax/match/protocol/step"
	"github.com/wdipax/match/state"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("creating bot: %s", err)
	}

	if v := os.Getenv("DEBUG"); v == "true" {
		bot.Debug = true
	}

	updateConfig := tgbotapi.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	s := state.New()

	admin := os.Getenv("ADMIN_USER_NAME")

	started := time.Now()

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		cancel()
	}()

	report := changeStageReporter()

	log.Println("started")

loop:
	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			break loop
		case update := <-updates:
			e := tgevent.New(update, admin, s.Step(), started)
			if e == nil {
				continue
			}

			r := s.Process(e)

			report(s.Step())

			if r == nil {
				continue
			}

			for _, m := range tgresponse.From(r, bot.Self.UserName, s.Step(), s.Admin()) {
				_, err := bot.Send(m)
				if err != nil {
					log.Printf("sending message: %s", err)
				}
			}
		}
	}

	if ctx.Err() != nil {
		log.Println("stopped normally")
	}
}

func changeStageReporter() func(current int) {
	var prev int

	return func(current int) {
		if prev != current {
			log.Printf("stage changed: %s -> %s", step.Name(prev), step.Name(current))

			prev = current
		}
	}
}
