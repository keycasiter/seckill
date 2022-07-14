package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"seckil/util"
)

var Conf *Config

type Config struct {
	MySQL `yaml:"mysql"`
}

type MySQL struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
}

func Init() {
	parseYamlConfig("config.yml")
}

func parseYamlConfig(path string) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		panic(fmt.Sprintf("parseYamlConfig err:%v", err))
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	log.Printf(fmt.Sprintf("config:\n%s\n", util.Marshal(conf)))
	Conf = conf
}
