package method

import (
	"context"
	"seckil/logic"
	"seckil/model/vo"
	"seckil/util"
)

type SecKillSrvImpl struct {
	Ctx context.Context
}

func NewSecKillSrvImpl(ctx context.Context) *SecKillSrvImpl {
	return &SecKillSrvImpl{Ctx: ctx}
}

// 查询商品信息
func (srv *SecKillSrvImpl) QueryGoodsInfo(ctx context.Context, req *vo.GoodsInfoQueryReq) *vo.GoodsInfoQueryResp {
	util.PrintLog(req)
	resp := logic.NewSecKillSrvLogicImpl(ctx).QueryGoodsInfo(req)
	util.PrintLog(resp)
	return resp
}

// 查询商品库存信息
func (srv *SecKillSrvImpl) QueryStockInfo(ctx context.Context, req *vo.StockInfoQueryReq) *vo.StockInfoQueryResp {
	util.PrintLog(req)
	resp := logic.NewSecKillSrvLogicImpl(ctx).QueryStockInfo(req)
	util.PrintLog(resp)
	return resp
}

// 查询商品、库存组合信息
func (srv *SecKillSrvImpl) QueryGoodsStockRichInfo(ctx context.Context, req *vo.GoodsStockRichInfoQueryReq) *vo.GoodsStockRichInfoQueryResp {
	util.PrintLog(req)
	resp := logic.NewSecKillSrvLogicImpl(ctx).QueryGoodsStockRichInfo(req)
	util.PrintLog(resp)
	return resp
}

// 秒杀下单
func (srv *SecKillSrvImpl) CreateOrder(ctx context.Context, req *vo.OrderCreateReq) *vo.OrderCreateResp {
	util.PrintLog(req)
	resp := logic.NewSecKillSrvLogicImpl(ctx).CreateOrder(req)
	util.PrintLog(resp)
	return resp
}
