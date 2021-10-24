package routes

import (
	"fog/common"
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
	Content []template.HTML
}

type TplEngine interface {
	RegisterTemplate(templatePath string) error
	Render(w http.ResponseWriter, page *Page)
}

type TemplatingEngine struct {
	logger              common.Logger
	registeredTemplates *template.Template
}

func NewTplEngine(logger common.Logger) *TemplatingEngine {
	return &TemplatingEngine{logger: logger}
}

func Render(w http.ResponseWriter, page *Page) {
	registeredTemplates.ExecuteTemplate(w, "main", page)
}

func (t *TemplatingEngine) RegisterTemplate(templatePath string) error {
	newTpl, err := t.registeredTemplates.ParseFiles(templatePath)
	if err == nil {
		t.registeredTemplates = newTpl
	}

	return err
}
