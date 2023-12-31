package global

import (
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var BasePath string
var Conf *Config
var DB *gorm.DB

type Config struct {
	Database string `yaml:"database"`
}

func Init() {
	LoadConfig()
	InitGormDB()
}
func LoadConfig() {
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

func InitGormDB() {
	var err error
	DB, _ = gorm.Open(mysql.Open(Conf.Database), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	println(DB.Config.Dialector.Name())

}
