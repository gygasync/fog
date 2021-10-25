package routes

import (
	"bytes"
	"fog/common"
	"fog/services"
	"fog/web/viewmodels"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type FilesRoute struct {
	logger           common.Logger
	tplEngine        TplEngine
	fileService      services.IFileService
	directoryService services.IDirectoryService

	internalTplEngine *template.Template
}

func NewFilesRoute(logger common.Logger, tplEngine TplEngine, fileService services.IFileService, directoryService services.IDirectoryService) *FilesRoute {
	return &FilesRoute{
		logger:            logger,
		tplEngine:         tplEngine,
		fileService:       fileService,
		directoryService:  directoryService,
		internalTplEngine: template.Must(template.ParseFiles("./web/static/templates/fileList.template.html"))}
}

func (i *FilesRoute) generateComponent(viewmodel *viewmodels.FilesInDirs) string {
	var buf bytes.Buffer
	i.internalTplEngine.ExecuteTemplate(&buf, "fileList", viewmodel)
	return buf.String()
}

func (i *FilesRoute) HandleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data, err := i.directoryService.GetChildren(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		i.logger.Warn("could not find files ", err)
		return
	}

	var bodyContent []template.HTML
	component := i.generateComponent(data)
	bodyContent = append(bodyContent, template.HTML(component))

	page := Page{Header: Header{Title: "Directories"}, Body: Body{Content: bodyContent}}
	i.tplEngine.Render(w, "main", &page)
}

func (i *FilesRoute) HandlePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
