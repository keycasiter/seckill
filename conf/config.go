package conf

import (
	"fmt"
	"github.com/ian-kent/go-log/log"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"path/filepath"
	"seckil/util"
)

var Conf *Config

type Config struct {
	MySQL    `yaml:"mysql"`
	Redis    `yaml:"redis"`
	RocketMq `yaml:"rocketmq"`
}

type MySQL struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

type RocketMq struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func Init() {
	rootPath := util.GetRootPath()
	confPath := path.Join(rootPath, "conf")
	path := filepath.Join(confPath, "config.yml")
	parseYamlConfig(path)
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
