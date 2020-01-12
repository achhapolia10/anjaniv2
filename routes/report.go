package routes

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/achhapolia10/inventory-manager/model"
	"github.com/julienschmidt/httprouter"
)

//GetDailyReport for route /product method GET
func GetDailyReport(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		t := template.Must(template.ParseGlob("views/components/navbar.comp"))
		t.ParseFiles("views/report.html")
		data := struct {
			U User
		}{
			currentUser,
		}
		t.ExecuteTemplate(w, "report.html", data)
	}
}

//PostDailyReport for route /product method POST
func PostDailyReport(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		d := req.FormValue("fdate")
		r := model.GetDailyReport(d)
		p, err := json.Marshal(r)
		if err != nil {
			log.Printf("Error in Marshalling Dialy Report : %v", err)
		}
		w.Write(p)
	}
}
