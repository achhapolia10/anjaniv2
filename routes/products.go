package routes

import (
	"html/template"
	"net/http"

	"github.com/achhapolia10/anjaniv2/opdatabase"

	"github.com/julienschmidt/httprouter"
)

//GetProducts Handler for route / method GET
func GetProducts(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/products.html")
	p, r := opdatabase.SelectProduct()
	if r {
		t.ExecuteTemplate(w, "products.html", p)
	}
}
