package web

import (
	"fog/web/routes"

	"github.com/julienschmidt/httprouter"
)

func New() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", routes.Index)

	return router
}
