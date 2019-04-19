package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//GetStock Handler for route / method GET
func GetStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/stock.html")

	t.ExecuteTemplate(w, "stock.html", "")
}
