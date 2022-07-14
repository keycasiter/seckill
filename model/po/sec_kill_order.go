package po

type SecKillOrder struct {
	Id         int64  `gorm:"primaryKey" json:"id" column:"id"`
	UserId     string `json:"user_id" column:"user_id"`
	SkuCode    string `json:"sku_code" column:"sku_code"`
	SkuNum     int64  `json:"sku_num" column:"sku_num"`
	CreateTime int64  `json:"create_time" column:"create_time"`
	UpdateTime int64  `json:"update_time" column:"update_time"`
	state      int64  `json:"state" column:"state"`
	OrderState int64  `json:"order_state" column:"order_state"`
}
