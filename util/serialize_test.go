package util

import (
	"fmt"
	"testing"
)

func TestJsonSerializer_Marshal(t *testing.T) {
	type User struct {
		Name string
		Age int
	}

	fmt.Println(Marshal(&User{Name: "zhangsan",Age: 20}))
}