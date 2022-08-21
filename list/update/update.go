package update

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type skuStruct struct {
	SKU []string `json:"sku"`
}

func Update(s interface{}) {
	file, err := os.Create("./json/list.json")
	result, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	fmt.Fprintln(mw, string(result))
}
