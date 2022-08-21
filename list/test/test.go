package listTest

import (
	"regexp"
)

type skuStruct struct {
	SKU []string `json:"sku"`
}

func IsValidSKU(s string) string {
	re := regexp.MustCompile("[0-9]+")
	c := re.FindAllString(s, -1)
	if !(c == nil) {
		return c[0]
	}
	return "nope"
}

func IsDupeSKU(l []string, s string) bool {
	for i := 0; i < len(l); i++ {
		if l[i] == s {
			return true
		}
	}
	return false
}

func IsListEmpty(l []string) bool {
	if len(l) == 0 {
		return true
	}
	return false
}
