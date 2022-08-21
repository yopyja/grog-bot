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

	if isSiteUp() {
		fmt.Println("Grog is alive 🍻")
		c := cron.New()
		c.AddFunc("@every 5m", stockBotFunc)
		c.Start()

		listBot.Run()
		<-make(chan struct{})
		return
	} else {
		fmt.Println("Grog is dead 😭")
	}
}

func stockBotFunc() {
	time := time.Now().Hour()
	if (time > 10) && (time < 18) {
		stockBot.Run()
		return
	}
}

func isSiteUp() bool {
	return stockTest.IsRealItem(config.TestSKU)
}
