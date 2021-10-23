package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// tmpl := template.Must(template.ParseFiles("./web/static/base.html"))
	// tmpl.Execute(w, nil)
	Render(w)
}
