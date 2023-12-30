package dao2

import (
	"mycity/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB2 *gorm.DB

func InitGormDB() {
	var err error
	DB2, _ = gorm.Open(mysql.Open(global.Conf.Database), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	println(DB2.Config.Dialector.Name())

}
