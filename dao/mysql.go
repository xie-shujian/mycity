package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	DB, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/world")

	DB.SetMaxIdleConns(2)
	fmt.Println("open mysql success")

}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("close mysql success")
	}
}
