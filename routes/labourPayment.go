package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//GetLabourPayment Handler for route / method GET
func GetLabourPayment(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/lp.html")

	t.ExecuteTemplate(w, "lp.html", "")
}
