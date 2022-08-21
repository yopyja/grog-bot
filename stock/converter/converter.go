package converter

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

func Run() DrinkInfo {
	jsonFile, err := os.Open("./json/temp.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	var drinks Drinks

	json.Unmarshal(byteValue, &drinks)

	return drinks.DrinkInfo[0]
}
