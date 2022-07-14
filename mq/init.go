package mq

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/ian-kent/go-log/log"
	"seckil/conf"
)

const (
	sec_kill_order_topic = "sec_kill_order"
)

var MqProducer rocketmq.Producer

func Init() {
	conf := conf.Conf.RocketMq
	endPoint := []string{fmt.Sprintf("%s:%s", conf.Host, conf.Port)}
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(endPoint),
		producer.WithNsResolver(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(2))
	if err != nil {
		panic(fmt.Sprintf("mq init err:%v", err))
	}
	err = p.Start()
	if err != nil {
		panic(fmt.Sprintf("mq producer start err:%v", err))
	}
	MqProducer = p
}

func SendSyncMessage(message string) error {
	result, err := MqProducer.SendSync(context.Background(), &primitive.Message{
		Topic: sec_kill_order_topic,
		Body:  []byte(message),
	})
	fmt.Println(result.Status, result)

	if err != nil {
		log.Fatalf("send mq err:%v", err)
		return err
	}
	return nil
}
