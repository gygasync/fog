package routes

import (
	"fog/common"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Route interface {
	HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	HandlePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type IndexRoute struct {
	logger    common.Logger
	tplEngine TplEngine
}

func NewIndexRoute(logger common.Logger, tplEngine TplEngine) *IndexRoute {
	return &IndexRoute{logger: logger, tplEngine: tplEngine}
}

func (i *IndexRoute) HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var bodyContent []template.HTML
	bodyContent = append(bodyContent, template.HTML("<h1>Welcome</h1>"))
	bodyContent = append(bodyContent, template.HTML("<p>Hope you are having a nice day</p>"))
	header := Header{Title: "FOG!"}
	i.tplEngine.Render(w, "main", &Page{Header: header, Body: Body{Content: bodyContent}})
}

func (i *IndexRoute) HandlePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
