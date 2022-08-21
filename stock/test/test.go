package stockTest

import (
	"grog-bot/stock/fetch"
	"os"
)

func IsRealItem(s string) bool {
	fetch.Item(s)

	tempFile, err := os.Stat("./json/temp.json")
	if err != nil {
	}

	if !(tempFile.Size() == 66) {
		return true
	}
	return false
}
