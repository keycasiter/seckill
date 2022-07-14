package seckill

import (
	"seckil/conf"
	"seckil/dal/mysql"
	"seckil/dal/redis"
	"seckil/mq"
)

func main() {
	//***初始化***
	//配置文件
	conf.Init()
	//mysql
	mysql.Init()
	//redis
	redis.Init()
	//mq
	mq.Init()
}
