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
	"github.com/wdipax/match/state"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("creating bot: %s", err)
	}

	// TODO: do we need this?
	bot.Debug = true

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
			if r == nil {
				continue
			}

			for _, m := range tgresponse.From(r, bot.Self.UserName, s.Step(), s.Admin()) {
				bot.Send(m)
			}

		}
	}

	if ctx.Err() != nil {
		log.Println("stopped normally")
	}
}
