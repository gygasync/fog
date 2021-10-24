package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// tmpl := template.Must(template.ParseFiles("./web/static/base.html"))
	// tmpl.Execute(w, nil)
	var bodyContent []template.HTML
	bodyContent = append(bodyContent, template.HTML("<h1>Welcome</h1>"))
	bodyContent = append(bodyContent, template.HTML("<p>Hope you are having a nice day</p>"))
	header := Header{Title: "FOG!"}
	Render(w, &Page{Header: header, Body: Body{Content: bodyContent}})
}
