package list

import (
	"encoding/json"
	"fmt"
	a "grog-bot/list/add"
	r "grog-bot/list/remove"
	"grog-bot/list/test"
	u "grog-bot/list/update"
	stockTest "grog-bot/stock/test"
	"io"
	"os"
)

type skuStruct struct {
	SKU []string `json:"sku"`
}

func ListManager(f int, s string) string {
	s = listTest.IsValidSKU(s)

	jsonFile, err := os.Open("./json/list.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var skuList skuStruct

	json.Unmarshal(byteValue, &skuList)

	// Adding for 0 : Removing for 1
	if stockTest.IsRealItem(s) {
		if f == 0 {
			if !(listTest.IsDupeSKU(skuList.SKU, s)) {
				skuList.SKU = a.Add(skuList.SKU, s)
				u.Update(interface{}(skuList))
				return "SKU: `" + s + "` has been added successfully"
			} else {
				return "SKU: `" + s + "` already exist"
			}
		} else {
			if listTest.IsDupeSKU(skuList.SKU, s) {
				skuList.SKU = r.Remove(skuList.SKU, s)
				u.Update(interface{}(skuList))
				return "SKU: `" + s + "` has been removed successfully"
			} else {
				return "SKU: `" + s + "` doesn't exist"
			}
		}
	} else {
		return "SKU: `" + s + "` does not exist in their database"
	}

}
