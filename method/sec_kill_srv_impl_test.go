package method

import (
	"context"
	"github.com/ian-kent/go-log/log"
	"seckil/conf"
	"seckil/dal/mysql"
	"seckil/dal/redis"
	"seckil/model/vo"
	"seckil/mq"
	"seckil/util"
	"testing"
)

func TestSecKillSrvImpl_CreateOrder(t *testing.T) {
	//配置文件
	conf.Init()
	//mysql
	mysql.Init()
	//redis
	redis.Init()
	//mq
	mq.Init()

	req := &vo.OrderCreateReq{
		UserId:  "zhangsan",
		SkuCode: "SKU001",
		SkuNum:  1,
	}
	resp := NewSecKillSrvImpl(context.Background()).CreateOrder(req)
	log.Printf("resp:%v", util.Marshal(resp))
}

func TestSecKillSrvImpl_QueryGoodsInfo(t *testing.T) {
	req := &vo.GoodsInfoQueryReq{
		SkuCode: "123138",
	}
	resp := NewSecKillSrvImpl(context.Background()).QueryGoodsInfo(req)
	log.Printf("resp:%v", resp)
}
