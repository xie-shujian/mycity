package controls3

import (
	"mycity/e"
	"mycity/s"

	"github.com/kataras/iris/v12"
)

func CityPage(ctx iris.Context) {
	page := ctx.URLParamIntDefault("page", 1)
	cites := s.PageCity((page-1)*20, 20)

	//ctx.ViewData("message", "hello world")
	ctx.ViewData("cities", cites)
	if err := ctx.View("cities.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}
func CityPage2(ctx iris.Context) {
	page := ctx.URLParamIntDefault("page", 1)
	pageCity := &e.PageCity{}
	pageCity.Page.PageSize = 20
	pageCity.Page.CurrentPage = page
	s.PageCity2(pageCity)

	//ctx.ViewData("message", "hello world")
	ctx.ViewData("pageCity", pageCity)
	if err := ctx.View("cities2.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}
func CityPageView(ctx iris.Context) {

	if err := ctx.View("cities3.html"); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}
func CityPageData(ctx iris.Context) {
	page := ctx.URLParamIntDefault("page", 1)
	pageCity := &e.PageCity{}
	pageCity.Page.PageSize = 20
	pageCity.Page.CurrentPage = page
	s.PageCity2(pageCity)
	ctx.JSON(pageCity)
}
