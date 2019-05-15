package routes

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/achhapolia10/anjaniv2/model"
	"github.com/julienschmidt/httprouter"
)

//GetLabourPayment Handler for route / method GET
func GetLabourPayment(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/navbar.comp"))
	t.ParseFiles("views/lp.html")

	t.ExecuteTemplate(w, "lp.html", "")
}

//PostLabourPayment Handler for route / method Post
func PostLabourPayment(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		log.Println("Error in PostLabourPayment: ", err)
	}
	fromDate := req.FormValue("from")
	toDate := req.FormValue("to")
	lp, _ := model.GetLabourPayment(fromDate, toDate)
	io.WriteString(w, lp)
}
