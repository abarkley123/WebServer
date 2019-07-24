package main

import (
	"html/template"
	"log"
    "net/http"
    "go.isomorphicgo.org/go/isokit"

    "github.com/WebServer/common"
	"github.com/gorilla/mux"
)

var WebAppRoot = "/Users/andrew/Desktop/git/go/src/github.com/WebServer";

// Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
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

func AboutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "./templates/about.html", nil)
	})   
}

func ContactHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "./templates/contact.html", nil)
	})   
}

func main() {
    env := common.Env{}
    isokit.TemplateFilesPath = WebAppRoot + "/templates"
	isokit.TemplateFileExtension = ".html"
	ts := isokit.NewTemplateSet()
	ts.GatherTemplates()
	env.TemplateSet = ts
    
    r := mux.NewRouter()
    r.Handle("/", MainHandler()).Methods("GET")
    r.Handle("/portfolio.html", PortfolioHandler()).Methods("GET")
    r.Handle("/about.html", AboutHandler()).Methods("GET")
    r.Handle("/contact.html", ContactHandler()).Methods("GET")

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", r)
    http.Handle("/portfolio.html", r)
    http.Handle("/about.html", r)
    http.Handle("/contact.html", r)
    http.Handle("/static/", http.StripPrefix("/static", fs))
    
    http.ListenAndServe(":8080", nil)
}