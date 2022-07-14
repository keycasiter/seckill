package mysql

import (
	"log"
	"seckil/model/po"
	"testing"
)

func TestInit(t *testing.T) {
	Init()

	order := &po.SecKillOrder{}
	if err := Db.Find(order).Error; err != nil {
		log.Printf("err:%v", err)
	}
	log.Printf("%v", order)
}
