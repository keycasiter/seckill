package seckill

import "seckil/model"

// 服务接口能力定义
type ISecKillSrv interface {
	// 查询商品信息
	QueryGoodsInfo(req *model.GoodsInfoQueryReq) *model.GoodsInfoQueryResp
	// 查询商品库存信息
	QueryStockInfo(req *model.StockInfoQueryReq) *model.StockInfoQueryResp
	// 查询商品、库存组合信息
	QueryGoodsStockRichInfo(req *model.GoodsStockRichInfoQueryReq) *model.GoodsStockRichInfoQueryResp
	// 秒杀下单
	CreateOrder(req *model.OrderCreateReq) *model.OrderCreateResp
}
