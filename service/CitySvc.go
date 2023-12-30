package service

import (
	"mycity/dao"
	"mycity/e"
)

func QueryAllCity() []e.City {
	var cities []e.City
	sql := "SELECT ID, Name, CountryCode, District, Population FROM city"
	rows, err := dao.DB.Query(sql)
	if err != nil {
		println(err.Error())
	}
	for rows.Next() {
		var city e.City
		rows.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)

		println(city.Name)
		cities = append(cities, city)
	}
	return cities
}
func QueryCity(id int64) *e.City {
	var city e.City
	sql := "SELECT ID, Name, CountryCode, District, Population FROM city where ID = ?"
	row := dao.DB.QueryRow(sql, id)
	row.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
	return &city
}
