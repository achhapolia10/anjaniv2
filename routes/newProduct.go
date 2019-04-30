package routes

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//GetNewProduct Handler for route / method GET
func GetNewProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/newProduct.html")

	t.ExecuteTemplate(w, "newProduct.html", "")
}
