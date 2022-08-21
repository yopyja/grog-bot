package grab

import (
	"encoding/json"
	"fmt"
	"grog-bot/list/test"
	"io"
	"os"
	"strconv"
)

type skuStruct struct {
	SKU []string `json:"sku"`
}

func GrabList() string {
	jsonFile, err := os.Open("./json/list.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var skuList skuStruct

	json.Unmarshal(byteValue, &skuList)

	if !(listTest.IsListEmpty(skuList.SKU)) {
		var message string

		message += "```c\n"

		for i := 0; i < len(skuList.SKU); i++ {
			if i < 9 {
				message += "0" + strconv.Itoa(i+1) + ":\t" + skuList.SKU[i] + "\n"
			} else {
				message += strconv.Itoa(i+1) + ":\t" + skuList.SKU[i] + "\n"
			}
		}

		message += "```"

		return message
	}

	return "```\nList Empty```"

}
