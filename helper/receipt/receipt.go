package receipt

import (
	"strconv"
	"time"
)

func GenerateReceipt() string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	receipt := "INV" + timestamp
	return receipt
}