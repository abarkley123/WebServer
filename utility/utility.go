package utility

import (
	"html/template"
	"log"
	"net/http"
)

var WebAppRoot = "/Users/andrew/Desktop/git/go/src/github.com/WebServer";

// Template rendering function
func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(WebAppRoot + "/templates/index.html", WebAppRoot + "/templates/header.html", WebAppRoot + "/templates/footer.html")
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func RenderGatedTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile, WebAppRoot+"/templates/gatedheader.html", WebAppRoot+"/templates/footer.html")
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}