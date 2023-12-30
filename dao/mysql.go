package dao

import (
	"database/sql"
	"fmt"
	"mycity/global"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	DB, _ = sql.Open("mysql", global.Conf.Database)

	DB.SetMaxIdleConns(2)
	fmt.Println("open mysql success")

}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("close mysql success")
	}
}
