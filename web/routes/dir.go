package routes

import (
	"bytes"
	"fog/common"
	"fog/db/genericmodels"
	"fog/services"
	"fog/work"
	"fog/work/definition"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type DirRoute struct {
	logger             common.Logger
	tplEngine          TplEngine
	service            services.IDirectoryService
	directoryWorkQueue work.IWorkQueue

	internalTplEngine *template.Template
}

func NewDirRoute(logger common.Logger, tplEngine TplEngine, service services.IDirectoryService, directoryWorkQueue work.IWorkQueue) *DirRoute {
	return &DirRoute{
		logger:             logger,
		tplEngine:          tplEngine,
		service:            service,
		directoryWorkQueue: directoryWorkQueue,
		internalTplEngine:  template.Must(template.ParseFiles("./web/static/templates/directoryList.template.html"))}
}

func (i *DirRoute) generateComponent(dirs []*genericmodels.Directory) string {
	var buf bytes.Buffer
	i.internalTplEngine.ExecuteTemplate(&buf, "directoryList", dirs)
	return buf.String()
}

func (i *DirRoute) HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := i.service.List(100, 0)
	var bodyContent []template.HTML
	component := i.generateComponent(data)
	bodyContent = append(bodyContent, template.HTML(component))

	page := Page{Header: Header{Title: "Directories"}, Body: Body{Content: bodyContent}}
	i.tplEngine.Render(w, "main", &page)
}

func (i *DirRoute) HandlePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	path := r.FormValue("path")
	if path != "" {
		work := definition.NewDirectoryWork(path)
		i.directoryWorkQueue.PostTask(work)
		http.Redirect(w, r, "dir", http.StatusFound)
	}
}
