package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//GetEntry Handler for route / method GET
func GetEntry(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/entry.html")

	t.ExecuteTemplate(w, "entry.html", "")
}
