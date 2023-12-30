package global

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database string `yaml:"database"`
}

var BasePath string
var Conf *Config

func Init() {
	BasePath, _ := os.Getwd()
	content, err := os.ReadFile(BasePath + "/resources/config.yml")
	if err != nil {
		println(err.Error())
	}
	err = yaml.Unmarshal(content, &Conf)
	if err != nil {
		println(err.Error())
	}
	println(Conf.Database)
}
