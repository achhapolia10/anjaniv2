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
	day1 := req.FormValue("day1")
	day2 := req.FormValue("day2")
	day3 := req.FormValue("day3")
	day4 := req.FormValue("day4")
	day5 := req.FormValue("day5")
	day6 := req.FormValue("day6")
	day7 := req.FormValue("day7")
	lp, _ := model.GetLabourPayment(day1, day2, day3, day4, day5, day6, day7)
	io.WriteString(w, lp)
}
