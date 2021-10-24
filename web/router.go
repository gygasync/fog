package web

import (
	"fmt"
	"fog/common"
	"fog/web/routes"

	"github.com/julienschmidt/httprouter"
)

type ITplRouter interface {
	RegisterRoute(path string, method Method, route routes.Route) error
	Router() *httprouter.Router
}

type TplRouter struct {
	logger    common.Logger
	tplEngine routes.TplEngine

	router *httprouter.Router
}

func NewTplRouter(logger common.Logger, tplEngine routes.TplEngine) *TplRouter {
	return &TplRouter{logger: logger, tplEngine: tplEngine, router: httprouter.New()}
}

func (r *TplRouter) RegisterRoute(path string, method Method, route routes.Route) error {
	switch method {
	case GET:
		r.router.GET(path, route.Handle)
	case POST:
		r.router.POST(path, route.Handle)
	default:
		return fmt.Errorf("failed registering route %s with method %s", path, method)
	}

	return nil
}

func (r *TplRouter) Router() *httprouter.Router {
	return r.router
}

// func New(logger common.Logger) *httprouter.Router {
// 	tplEngine := routes.NewTplEngine(logger)
// 	tplEngine.RegisterTemplate("./web/static/templates/main.template.html", "main")
// 	tplEngine.RegisterTemplate("./web/static/templates/body.template.html", "body")
// 	tplEngine.RegisterTemplate("./web/static/templates/header.template.html", "header")

// 	router := httprouter.New()

// 	router.GET("/", routes.Index)

// 	return router
// }
