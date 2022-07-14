package po

type SecKillGoods struct {
	Id         int64  `gorm:"primaryKey" json:"id" column:"id"`
	SkuCode    string `json:"sku_code" column:"sku_code"`
	Stock      int64  `json:"stock" column:"stock"`
	CreateTime int64  `json:"create_time" column:"create_time"`
	UpdateTime int64  `json:"update_time" column:"update_time"`
	state      int64  `json:"state" column:"state"`
}
