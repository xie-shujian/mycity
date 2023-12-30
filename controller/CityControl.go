package controller

import (
	"html/template"
	"mycity/service"
	"net/http"
	"strconv"
)

func CityController(w http.ResponseWriter, r *http.Request) {
	sid := r.URL.Path[len("/city/"):]
	id, _ := strconv.ParseInt(sid, 10, 64)
	city := service.QueryCity(id)
	t, err := template.ParseFiles("viewer/city.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, city)
}

func CitiesController(w http.ResponseWriter, r *http.Request) {
	cities := service.QueryAllCity()

	t, err := template.ParseFiles("viewer/cities.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, cities)
}

func CitiesControllerT(w http.ResponseWriter, r *http.Request) {
	//cities := service.QueryAllCity()
	cites := []string{"a", "b", "c", "d", "e", "f"}
	t, err := template.ParseFiles("viewer/cities.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, cites)
}
