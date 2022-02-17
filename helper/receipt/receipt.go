package receipt

import (
	"strconv"
	"time"
)

func GenerateReceipt(id string) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	receipt := "INV" + timestamp + id
	return receipt
}