package logic

import (
	"context"
	"seckil/model/vo"
)

type SecKillSrvLogic interface {
	// 查询商品信息
	QueryGoodsInfo(req *vo.GoodsInfoQueryReq) (*vo.GoodsInfoQueryResp, error)
	// 查询商品库存信息
	QueryStockInfo(req *vo.StockInfoQueryReq) (*vo.StockInfoQueryResp, error)
	// 查询商品、库存组合信息
	QueryGoodsStockRichInfo(req *vo.GoodsStockRichInfoQueryReq) (*vo.GoodsStockRichInfoQueryResp, error)
	// 秒杀下单
	CreateOrder(req *vo.OrderCreateReq) (*vo.OrderCreateResp, error)
}

type SecKillSrvLogicImpl struct {
	Ctx context.Context
}

func NewSecKillSrvLogicImpl(ctx context.Context) *SecKillSrvLogicImpl {
	return &SecKillSrvLogicImpl{Ctx: ctx}
}

// 查询商品信息
func (srv *SecKillSrvLogicImpl) QueryGoodsInfo(req *vo.GoodsInfoQueryReq) *vo.GoodsInfoQueryResp {
	return nil
}

// 查询商品库存信息
func (srv *SecKillSrvLogicImpl) QueryStockInfo(req *vo.StockInfoQueryReq) *vo.StockInfoQueryResp {
	return nil
}

// 查询商品、库存组合信息
func (srv *SecKillSrvLogicImpl) QueryGoodsStockRichInfo(req *vo.GoodsStockRichInfoQueryReq) *vo.GoodsStockRichInfoQueryResp {
	return nil
}

// 秒杀下单
func (srv *SecKillSrvLogicImpl) CreateOrder(req *vo.OrderCreateReq) *vo.OrderCreateResp {
	return nil
}
