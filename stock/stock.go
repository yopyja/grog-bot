package stock

import (
	"encoding/json"
	"fmt"
	"grog-bot/config"
	conv "grog-bot/stock/converter"
	"grog-bot/stock/fetch"
	"io"
	"log"
	"os"
	"strconv"
)

type Drinks struct {
	DrinkInfo []DrinkInfo `json:"data"`
}

type DrinkInfo struct {
	SKU          string  `json:"sku"`
	Name         string  `json:"name"`
	WarehouseQty int     `json:"warehouseQty"`
	StoreQty     int     `json:"storeQty"`
	CurrentPrice float64 `json:"currentPrice"`
}

type skuStruct struct {
	SKU []string `json:"sku"`
}

type newStockMessage struct {
	Message []string `json:"message"`
}

func StockRunner() {
	// Load old stock
	jsonFile, err := os.Open("./json/store.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := io.ReadAll(jsonFile)
	var oldStock Drinks
	json.Unmarshal(byteValue, &oldStock)

	// Load files
	storeFile, err := os.Create("./json/store.json")
	listFile, err := os.Open("./json/list.json")
	messageFile, err := os.Create("./json/message.json")

	// Error Handling
	if err != nil {
		fmt.Println(err)
	}

	//Load latest sku list file
	byteValueList, _ := io.ReadAll(listFile)
	var skuList skuStruct
	json.Unmarshal(byteValueList, &skuList)

	// Load latest stock into this var
	var newStock Drinks

	// Go through sku list and add stock to newStock
	for i := 0; i < len(skuList.SKU); i++ {
		fetch.Item(skuList.SKU[i])
		newStock.DrinkInfo = append(newStock.DrinkInfo, DrinkInfo(conv.Run()))
	}

	// If new stock store in the following
	var newMessage newStockMessage

	//Compare two structs look for new stock increase
	for n := 0; n < len(newStock.DrinkInfo); n++ {
		for o := 0; o < len(oldStock.DrinkInfo); o++ {
			if newStock.DrinkInfo[n].SKU == oldStock.DrinkInfo[o].SKU {
				if newStock.DrinkInfo[n].StoreQty > oldStock.DrinkInfo[o].StoreQty {
					newMessage.Message = append(newMessage.Message, newStock.DrinkInfo[n].Name+"||StoreQty:\t"+strconv.Itoa(newStock.DrinkInfo[n].StoreQty)+"||"+config.ItemURL+newStock.DrinkInfo[n].SKU)
				}
			}
		}
	}

	// Save new stock to store.json
	result, err := json.Marshal(newStock)
	if err != nil {
		log.Println(err)
	}
	mw := io.MultiWriter(os.Stdout, storeFile)
	fmt.Fprintln(mw, string(result))

	//Save message array
	messageResult, err := json.Marshal(newMessage)
	mwMessage := io.MultiWriter(os.Stdout, messageFile)
	fmt.Fprintln(mwMessage, string(messageResult))

}
