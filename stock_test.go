package seckill

import (
	"fmt"
	"github.com/ian-kent/go-log/log"
	"seckil/conf"
	"seckil/dal/redis"
	"testing"
)

// 设置商品单一库存
func TestPutSkuSingleStock(t *testing.T) {
	skuCode := "SKU001"
	conf.Init()
	redis.Init()
	resp, err := redis.RedisClient.Incr(fmt.Sprintf("KEY:STOCK:%s", skuCode)).Result()
	if err != nil {
		log.Error("err:%v", err)
	}
	log.Info("resp:%v", resp)
}

// 设置商品多库存
func TestPutSkuMultiStock(t *testing.T) {
	skuCode := "SKU001"
	conf.Init()
	redis.Init()
	resp, err := redis.RedisClient.Incr(fmt.Sprintf("KEY:STOCK:%s", skuCode)).Result()
	if err != nil {
		log.Error("err:%v", err)
	}
	log.Info("resp:%v", resp)
}

// 清空商品库存
func TestClearSkuStock(t *testing.T) {

}
