package main

import (
	"mycity/control2"
	"mycity/controls3"
	"mycity/global"
	"net/http"

	"github.com/kataras/iris/v12"
)

func main() {
	global.Init()
	main4()
}

func main4() {
	//global.InitGormDB()
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.HandleDir("/statics", "./statics")
	//app.Get("/city/view", controls3.CityPageView)
	app.Get("/city/data", controls3.CityPageData)
	app.Listen(":8080")
}
func main3() {
	global.InitGormDB()
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))
	app.Get("/cities", controls3.CityPage2)
	app.Listen(":8080")
}

func main2() {
	global.InitGormDB()
	// var cities []entity.City

	// dao2.DB2.Limit(10).Offset(10).Find(&cities)
	// println(cities[0].CountryCode)
	http.HandleFunc("/cities", control2.PageCity)
	http.ListenAndServe(":8080", nil)
}

// func main1() {
// 	dao.InitDB()
// 	// http.HandleFunc("/city/", controller.CityController)
// 	//
// 	http.HandleFunc("/cities", controller.CitiesController)
// 	http.ListenAndServe(":8080", nil)
// 	dao.CloseDB()
// }
