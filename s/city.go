package s

import (
	"mycity/dao2"
	"mycity/e"
)

func PageCity(offset int, limit int) []e.City {
	return dao2.PageCity(offset, limit)
}

func PageCity2(pageCity *e.PageCity) {

	if pageCity.Page.TotalRecords == 0 {
		pageCity.Page.TotalRecords = dao2.CityTotal()
		pageCity.Page.TotalPages = (pageCity.Page.TotalRecords + pageCity.Page.PageSize - 1) / pageCity.Page.PageSize
	}

	limit := pageCity.Page.PageSize
	offset := (pageCity.Page.CurrentPage - 1) * pageCity.Page.PageSize

	pageCity.Cities = dao2.PageCity(offset, limit)

}
