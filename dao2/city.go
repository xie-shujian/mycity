package dao2

import (
	"mycity/e"
	"mycity/global"
)

func PageCity(offset int, limit int) []e.City {

	var cities []e.City

	global.DB.Limit(limit).Offset(offset).Find(&cities)
	return cities
}

func CityTotal() int {
	var cities []e.City
	result := global.DB.Find(&cities)
	return int(result.RowsAffected)
}
