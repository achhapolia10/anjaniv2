package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/achhapolia10/anjaniv2/opdatabase"

	"github.com/julienschmidt/httprouter"
)

//GetEntry Handler for route / method GET
func GetEntry(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	p, res := opdatabase.SelectProduct()
	if !res {
		fmt.Print("Error in querying all products in Entries")
	} else {
		t := template.Must(template.ParseGlob("views/components/*.comp"))
		t.ParseFiles("views/entry.html")
		t.ExecuteTemplate(w, "entry.html", p)
	}
}
