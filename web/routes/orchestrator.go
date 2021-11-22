package routes

import (
	"bytes"
	"fog/common"
	"fog/tasks"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrchestratorRoute struct {
	logger            common.Logger
	tplEngine         TplEngine
	internalTplEngine *template.Template
	orchestrator      tasks.IOrchestrator
}

func NewOrchestratorRoute(logger common.Logger, tplEngine TplEngine, orchestrator tasks.IOrchestrator) *OrchestratorRoute {
	return &OrchestratorRoute{
		logger:            logger,
		tplEngine:         tplEngine,
		orchestrator:      orchestrator,
		internalTplEngine: template.Must(template.ParseFiles("./web/static/templates/orchestratorDetails.template.html")),
	}
}

func (i *OrchestratorRoute) generateComponent(details []string) string {
	var buf bytes.Buffer
	i.internalTplEngine.ExecuteTemplate(&buf, "orchestratorDetails", details)
	return buf.String()
}

func (i *OrchestratorRoute) HandleGet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var bodyContent []template.HTML
	component := i.generateComponent(i.orchestrator.GetDetails())
	bodyContent = append(bodyContent, template.HTML(component))

	page := Page{Header: Header{Title: "Orchestrator"}, Body: Body{Content: bodyContent}}
	i.tplEngine.Render(w, "main", &page)
}

func (i *OrchestratorRoute) HandlePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusMethodNotAllowed)
}
