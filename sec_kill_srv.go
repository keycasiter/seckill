package seckill

import (
	"context"
	"seckil/model/vo"
)

// 服务接口能力定义
type ISecKillSrv interface {
	// 查询商品信息
	QueryGoodsInfo(ctx context.Context, req *vo.GoodsInfoQueryReq) *vo.GoodsInfoQueryResp
	// 秒杀下单
	CreateOrder(ctx context.Context, req *vo.OrderCreateReq) *vo.OrderCreateResp
}
