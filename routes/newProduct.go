package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/achhapolia10/anjaniv2/opdatabase"

	"github.com/achhapolia10/anjaniv2/model"

	"github.com/julienschmidt/httprouter"
)

//GetNewProduct Handler for route / method GET
func GetNewProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/newProduct.html")

	t.ExecuteTemplate(w, "newProduct.html", "")
}

//PostNewProduct creates a new Product
func PostNewProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}
	var n string
	var p, b, op, ob int
	var pr float64
	n = req.FormValue("name")
	p, err = strconv.Atoi(req.FormValue("box"))
	b, err = strconv.Atoi(req.FormValue("packet"))
	op, err = strconv.Atoi(req.FormValue("opackets"))
	ob, err = strconv.Atoi(req.FormValue("oboxes"))
	pr, err = strconv.ParseFloat(req.FormValue("price"), 32)
	product := model.Product{
		0, n, p, b, pr, ob, op,
	}
	opdatabase.AddProduct(product)
	http.Redirect(w, req, "/products", 301)
}
