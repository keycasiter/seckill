package mysql

import (
	"github.com/ian-kent/go-log/log"
	"seckil/conf"
	"seckil/model/po"
	"testing"
)

func TestInit(t *testing.T) {
	conf.Init()
	Init()

	order := &po.SecKillOrder{}
	if err := Db.Find(order).Error; err != nil {
		log.Printf("err:%v", err)
	}
	log.Printf("%v", order)
}
