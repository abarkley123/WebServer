package handlers

import(
	"github.com/WebServer/utility"
	"net/http"
)

var WebAppRoot = "/Users/andrew/Desktop/git/go/src/github.com/WebServer";
                                 
func MainHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utility.RenderTemplate(w, "/Users/andrew/Desktop/git/go/src/github.com/WebServer/templates/index.html", nil)
	})
}