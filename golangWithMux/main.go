package main

import (
	"golang_with_mux/utilities"
	"html/template"
	"net/http"
)

type ComponentContext struct {
	error
}

func HomeHandler(t *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "home.html", nil)
	}
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := utilities.GetInfoLogger()
		log.Println(r.Method)
		h(w, r)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	allTemplates, err := template.ParseGlob("./templates/*.html")
	utilities.HandleError(err)
	http.HandleFunc("/", log(HomeHandler(allTemplates)))

	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: nil,
	}
	server.ListenAndServe()
}
