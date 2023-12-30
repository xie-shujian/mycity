package dao2

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB2 *gorm.DB

func InitGormDB() {
	var err error
	DB2, _ = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/world"), &gorm.Config{
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
