package fetch

import (
	"fmt"
	"grog-bot/config"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Item(sku string) {
	file, err := os.Create("./json/temp.json")

	url := config.URL
	method := "POST"

	payload := strings.NewReader(config.Payload[0] + sku + config.Payload[1])
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "1932")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Host", config.Host)
	req.Header.Add("Origin", config.Origin)
	req.Header.Add("Referer", config.Referer)

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	fmt.Fprintln(mw, string(body))
}
