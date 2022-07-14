package vo

type OrderCreateReq struct {
	//用户ID
	UserId string
	//购物商品SKU
	SkuCode string
	//购物商品梳理
	SkuNum int
}
