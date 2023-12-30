package control2

import (
	"html/template"
	"mycity/s"
	"net/http"
)

func PageCity(w http.ResponseWriter, r *http.Request) {
	cites := s.PageCity(10, 10)
	t, err := template.ParseFiles("viewer/cities.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, cites)
}
