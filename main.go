package main

import (
	"html/template"
	"log"
	"net/http"
    
    "github.com/gorilla/mux"
)

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func MainHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "./templates/index.html", nil)
	})
}

func PortfolioHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "./templates/portfolio.html", nil)
	})
}

func main() {
    r := mux.NewRouter()
    r.Handle("/", MainHandler()).Methods("GET")
    r.Handle("/portfolio.html", PortfolioHandler()).Methods("GET")

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", r)
    http.Handle("/portfolio.html", r)
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}