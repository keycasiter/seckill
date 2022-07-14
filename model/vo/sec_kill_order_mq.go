package vo

import "time"

type SecKillOrderMq struct {
	UserId    string
	SkuCode   string
	SkuNum    int
	OrderNo   string
	OrderTime time.Time
}
