package mq

import (
	"seckil/conf"
	"testing"
)

func TestSend(t *testing.T) {
	conf.Init()
	Init()

	SendSyncMessage("hello world")
}
