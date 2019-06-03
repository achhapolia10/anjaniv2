package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"

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
	lp := model.GetLabourPayment(day1, day2, day3, day4, day5, day6, day7)
	var response Response
	if len(lp) == 0 {
		response = Response{501, lp}
	} else {
		response = Response{301, lp}
	}
	p, e := json.Marshal(response)
	if e != nil {
		log.Println(e)
	}
	io.WriteString(w, string(p))
}

//S f
type S struct {
	Lps          []model.LPEntry
	Days         []string
	TotalAmmount float64
}

//GetPrintLabourPayment handler form route /print method Post
func GetPrintLabourPayment(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	url := req.URL
	q := url.Query()
	day1 := q.Get("day1")
	day2 := q.Get("day2")
	day3 := q.Get("day3")
	day4 := q.Get("day4")
	day5 := q.Get("day5")
	day6 := q.Get("day6")
	day7 := q.Get("day7")
	lp := model.GetLabourPayment(day1, day2, day3, day4, day5, day6, day7)
	t := template.Must(template.ParseGlob("views/components/navbar.comp"))
	t.Funcs(template.FuncMap{
		"shouldPrint": ShouldPrintL,
		"countDays":   CountDays,
		"roundMoney":  RoundMoney,
	}).ParseFiles("views/lpPrint.html")
	sort.Sort(lp)
	err := t.ExecuteTemplate(w, "lpPrint.html", S{Lps: lp, Days: []string{day1, day2, day3, day4, day5, day6, day7}, TotalAmmount: lp.TotalAmmount()})
	if err != nil {

		log.Printf("Error in Printining Labour Payment: %v", err)
	}
}

//ShouldPrintL should be printed or not
func ShouldPrintL(a int) bool {
	return a > 0
}

//CountDays Return Count Days to
func CountDays(a model.LPEntry) int {
	n := 0
	if a.Day1 > 0 {
		n++
	}
	if a.Day2 > 0 {
		n++
	}
	if a.Day3 > 0 {
		n++
	}
	if a.Day4 > 0 {
		n++
	}
	if a.Day5 > 0 {
		n++
	}
	if a.Day6 > 0 {
		n++
	}
	if a.Day7 > 0 {
		n++
	}
	return n
}

//RoundMoney rounds
func RoundMoney(a float64) string {
	return fmt.Sprintf("%.00f", a)
}
