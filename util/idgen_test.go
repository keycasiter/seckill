package util

import (
	"fmt"
	"testing"
)

func TestGenUuid(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GenUuid())
	}
}
