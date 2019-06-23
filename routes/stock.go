package routes

import (
	"encoding/json"
	"html/template"
	"log"
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
	b, err := json.Marshal(stockData)
	if err != nil {
		log.Printf("Error in Marshalling stock data: %v", err)
		return
	}
	w.Write(b)
}
