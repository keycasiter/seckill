package util

import (
	"context"
	"encoding/json"
	"github.com/ian-kent/go-log/log"
	"reflect"
)

var serializer Serializer

func init() {
	serializer = NewJsonSerializer(context.Background())
}

func Marshal(obj interface{}) string {
	return serializer.Marshal(obj)
}

func Unmarshal(str string, target interface{}) {
	//check
	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		panic("JsonSerializer Unmarshal target must be pointer obj")
	}

	serializer.Unmarshal(str, target)
}

type Serializer interface {
	//序列化
	Marshal(interface{}) string
	//反序列化
	Unmarshal(str string, target interface{})
}

type JsonSerializer struct {
	Ctx context.Context
}

func NewJsonSerializer(ctx context.Context) *JsonSerializer {
	return &JsonSerializer{Ctx: ctx}
}

func (s *JsonSerializer) Marshal(obj interface{}) string {
	var b []byte
	var err error
	if b, err = json.Marshal(obj); err != nil {
		log.Printf("JsonSerializer Marshal err:%v", err)
		return ""
	}
	return string(b)
}

func (s *JsonSerializer) Unmarshal(str string, target interface{}) {
	if err := json.Unmarshal([]byte(str), target); err != nil {
		log.Printf("JsonSerializer Unmarshal err:%v", err)
	}
}
