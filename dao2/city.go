package dao2

import "mycity/e"

func PageCity(offset int, limit int) []e.City {

	var cities []e.City

	DB2.Limit(limit).Offset(offset).Find(&cities)
	return cities
}

func CityTotal() int {
	var cities []e.City
	result := DB2.Find(&cities)
	return int(result.RowsAffected)
}
