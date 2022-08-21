package main

import (
	"github.com/robfig/cron/v3"
	listBot "grog-bot/bot/list"
	stockBot "grog-bot/bot/stock"
	"grog-bot/config"
	"log"
	"time"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	c := cron.New()
	c.AddFunc("@every 5m", stockBotFunc)
	c.Start()

	listBot.Run()
	<-make(chan struct{})
	return

}

func stockBotFunc() {
	time := time.Now().Hour()
	if (time > 10) && (time < 18) {
		stockBot.Run()
		return
	}
}
