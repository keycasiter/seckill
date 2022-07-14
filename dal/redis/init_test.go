package redis

import (
	"fmt"
	"seckil/conf"
	"testing"
)

func TestInit(t *testing.T) {
	conf.Init()
	Init()

	fmt.Println(redisClient.Ping().Result())
}
