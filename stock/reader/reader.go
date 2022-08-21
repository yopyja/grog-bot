package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

func OldStockGrabber() Drinks {
	jsonFile, err := os.Open("./json/store.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var oldStock Drinks

	json.Unmarshal(byteValue, &oldStock)

	return oldStock
}
