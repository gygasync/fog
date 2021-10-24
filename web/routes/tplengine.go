package routes

import (
	"fmt"
	"fog/common"
	"html/template"
	"net/http"
)

// Templating engine

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
	RegisterTemplate(templatePath string, templateName string) error
	Render(w http.ResponseWriter, templateName string, page *Page) error
}

type TemplatingEngine struct {
	logger              common.Logger
	registeredTemplates *template.Template
	templateEntries     []string
}

func NewTplEngine(logger common.Logger) *TemplatingEngine {
	return &TemplatingEngine{logger: logger}
}

// No need for speed here
func contains(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}

	return false
}

func (t *TemplatingEngine) Render(w http.ResponseWriter, templateName string, page *Page) error {
	if !contains(t.templateEntries, templateName) {
		return fmt.Errorf("template %s has not been registered ", templateName)
	}

	return t.registeredTemplates.ExecuteTemplate(w, templateName, page)
}

func (t *TemplatingEngine) RegisterTemplate(templatePath string, templateName string) error {
	newTpl, err := t.registeredTemplates.ParseFiles(templatePath)
	if err == nil {
		t.registeredTemplates = newTpl
		t.templateEntries = append(t.templateEntries, templateName)
	}

	return err
}
