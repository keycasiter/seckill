package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"seckil/conf"
	"time"
)

var RedisClient *redis.Client

func Init() {
	redisConf := conf.Conf.Redis
	opts := redis.Options{
		Addr:        redisConf.Host,     //redis地址
		Password:    redisConf.Password, //密码
		DialTimeout: time.Second * 5,    //超时时间
	}
	RedisClient = redis.NewClient(&opts)             //创建连接
	if err := RedisClient.Ping().Err(); err != nil { //心跳测试
		panic(fmt.Sprintf("redis client err:%v", err))
	}
}
