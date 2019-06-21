package routes

import (
	"html/template"
	"net/http"

	"github.com/achhapolia10/anjaniv2/model"
	"github.com/julienschmidt/httprouter"
)

//GetStock Handler for route / method GET
func GetStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/navbar.comp"))
	t.ParseFiles("views/stock.html")

	t.ExecuteTemplate(w, "stock.html", "")
}

//PostStock Handler for route / method POST query: fdate,tdate
func PostStock(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fromDate := req.FormValue("fdate")
	toDate := req.FormValue("tdate")
	stockData := model.AllStock(fromDate, toDate)
	w.Write([]byte(stockData))
}
