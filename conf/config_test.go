package conf

import (
	"log"
	"testing"
)

func TestParseYamlConfig(t *testing.T) {
	Init()

	log.Printf(Conf.UserName)
}
