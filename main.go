package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	listBot "grog-bot/bot/list"
	stockBot "grog-bot/bot/stock"
	"grog-bot/config"
	stockTest "grog-bot/stock/test"
	"log"
	"time"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	if healthCheck() {
		fmt.Println("Health Check Passed 🍻")
		c := cron.New()
		c.AddFunc("@every 5m", stockBotFunc)
		c.Start()

		listBot.Run()
		<-make(chan struct{})
		return
	} else {
		fmt.Println("Health Check Failed 💀")
	}
}

func stockBotFunc() {
	hour := time.Now().Hour()
	weekday := time.Now().Weekday()
	if weekday != 0 {
		if (hour >= 10) && (hour < 18) {
			stockBot.Run()
			return
		}
	}
}

func healthCheck() bool {
	return stockTest.IsRealItem(config.TestSKU)
}
