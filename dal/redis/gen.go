package redis

import "fmt"

const (
	StockKey = "KEY:STOCK:%s"
)

func GenStockKey(skuCode string) string {
	return fmt.Sprintf(StockKey, skuCode)
}
