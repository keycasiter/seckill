package method

import (
	"context"
	"seckil/model"
)

type SecKillSrvImpl struct {
	Ctx context.Context
}

func NewSecKillSrvImpl(ctx context.Context) *SecKillSrvImpl {
	return &SecKillSrvImpl{Ctx: ctx}
}

// 查询商品信息
func (srv *SecKillSrvImpl) QueryGoodsInfo(req *model.GoodsInfoQueryReq) *model.GoodsInfoQueryResp {
	return nil
}

// 查询商品库存信息
func (srv *SecKillSrvImpl) QueryStockInfo(req *model.StockInfoQueryReq) *model.StockInfoQueryResp {
	return nil
}

// 查询商品、库存组合信息
func (srv *SecKillSrvImpl) QueryGoodsStockRichInfo(req *model.GoodsStockRichInfoQueryReq) *model.GoodsStockRichInfoQueryResp {
	return nil
}

// 秒杀下单
func (srv *SecKillSrvImpl) CreateOrder(req *model.OrderCreateReq) *model.OrderCreateResp {
	return nil
}
