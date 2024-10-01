package helpers

import (
	"fmt"
	"strconv"
)

func ParsePrice(price string) float32 {
	floatValue, err := strconv.ParseFloat(price, 32)
	if err != nil {
		fmt.Println("Error parsing product price:", err)
		return 0
	}

	return float32(floatValue)
}
