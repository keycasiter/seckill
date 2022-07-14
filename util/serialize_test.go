package util

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestJsonSerializer_Marshal(t *testing.T) {
	fmt.Println(Marshal(&User{Name: "zhangsan", Age: 20}))
}

func TestJsonSerializer_Unmarshal(t *testing.T) {
	u := &User{}
	Unmarshal("{\"Name\":\"zhangsan\",\"Age\":20}", u)
	fmt.Println(u)
}
