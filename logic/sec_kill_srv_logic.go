package logic

import (
	"context"
	"github.com/ian-kent/go-log/log"
	"seckil/dal/redis"
	"seckil/model/vo"
	"seckil/mq"
	"seckil/rpc"
	"seckil/util"
	"time"
)

type SecKillSrvLogic interface {
	// 查询商品信息
	QueryGoodsInfo(req *vo.GoodsInfoQueryReq) (*vo.GoodsInfoQueryResp, error)
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

// 秒杀下单
func (srv *SecKillSrvLogicImpl) CreateOrder(req *vo.OrderCreateReq) *vo.OrderCreateResp {
	//1. 风控过滤
	if !rpc.ProcessRisk(req.UserId) {
		return &vo.OrderCreateResp{Result: vo.BuildRiskRefuseResult()}
	}
	//2. 库存扣减
	key := redis.GenStockKey(req.SkuCode)
	res, err := redis.RedisClient.Decr(key).Result()
	if err != nil {
		log.Fatalf("CreateOrder err:%v", err)
		return &vo.OrderCreateResp{Result: vo.BuildFailResult()}
	}
	if res < 0 {
		log.Println("CreateOrder stock not enough")
		return &vo.OrderCreateResp{Result: vo.BuildStockNotEnoughResult()}
	}
	//3. 发送下单MQ
	err = mq.SendSyncMessage(util.Marshal(&vo.SecKillOrderMq{
		UserId:    req.UserId,
		SkuCode:   req.SkuCode,
		SkuNum:    req.SkuNum,
		OrderNo:   util.GenUuid(),
		OrderTime: time.Time{},
	}))
	if err != nil {
		log.Fatalf("mq send err:%v", err)
		return &vo.OrderCreateResp{Result: vo.BuildFailResult()}
	}

	return &vo.OrderCreateResp{Result: vo.BuildSuccessResult()}
}
