package routes

import (
	"bytes"
	"fog/common"
	"fog/db/models"
	"fog/services"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DirRoute struct {
	logger    common.Logger
	tplEngine TplEngine
	service   services.IDirectoryService

	internalTplEngine *template.Template
}

func NewDirRoute(logger common.Logger, tplEngine TplEngine, service services.IDirectoryService) *DirRoute {
	return &DirRoute{
		logger:            logger,
		tplEngine:         tplEngine,
		service:           service,
		internalTplEngine: template.Must(template.ParseFiles("./web/static/templates/directoryList.template.html"))}
}

func (i *DirRoute) generateComponent(dirs []models.Directory) string {
	var buf bytes.Buffer
	i.internalTplEngine.ExecuteTemplate(&buf, "directoryList", dirs)
	return buf.String()
}

func (i *DirRoute) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := i.service.List(100, 0)
	var bodyContent []template.HTML
	component := i.generateComponent(data)
	bodyContent = append(bodyContent, template.HTML(component))

	page := Page{Header: Header{Title: "Directories"}, Body: Body{Content: bodyContent}}
	i.tplEngine.Render(w, "main", &page)
}
