package mq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"seckil/conf"
	"testing"
)

func Test_CreateTopic(t *testing.T) {
	conf.Init()
	conf := conf.Conf.RocketMq
	// 创建主题
	topicAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver([]string{fmt.Sprintf("%s:%s", conf.Host, conf.Port)})))
	if err != nil {
		fmt.Println(err)
	}
	err = topicAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate("sec_kill_order"),
		admin.WithBrokerAddrCreate(fmt.Sprintf("%s:%s", "127.0.0.1", "10911")),
	)
	fmt.Println(err)
}
