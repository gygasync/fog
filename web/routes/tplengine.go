package routes

import (
	"html/template"
	"net/http"
)

// Templating engine
var registeredTemplates = template.Must(template.ParseFiles(
	"./web/static/templates/main.template.html",
	"./web/static/templates/body.template.html",
	"./web/static/templates/header.template.html",
))

type Page struct {
	Header Header
	Body   Body
}

type Header struct {
	Title string
}

type Body struct {
	Content template.HTML
}

func Render(w http.ResponseWriter) {
	registeredTemplates.ExecuteTemplate(w, "main", &Page{Header: Header{Title: "FOG!"}, Body: Body{Content: template.HTML("<h1>Welcome</h1>")}})
}
