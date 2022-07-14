package conf

import (
	"github.com/ian-kent/go-log/log"
	"testing"
)

func TestParseYamlConfig(t *testing.T) {
	Init()

	log.Printf(Conf.UserName)
}
