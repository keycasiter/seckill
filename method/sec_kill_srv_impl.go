package method

import (
	"context"
	"seckil/logic"
	"seckil/model/vo"
)

type SecKillSrvImpl struct {
	Ctx context.Context
}

func NewSecKillSrvImpl(ctx context.Context) *SecKillSrvImpl {
	return &SecKillSrvImpl{Ctx: ctx}
}

// 查询商品信息
func (srv *SecKillSrvImpl) QueryGoodsInfo(req *vo.GoodsInfoQueryReq) *vo.GoodsInfoQueryResp {
	resp := logic.NewSecKillSrvLogicImpl(srv.Ctx).QueryGoodsInfo(req)
	return resp
}

// 秒杀下单
func (srv *SecKillSrvImpl) CreateOrder(req *vo.OrderCreateReq) *vo.OrderCreateResp {
	resp := logic.NewSecKillSrvLogicImpl(srv.Ctx).CreateOrder(req)
	return resp
}
