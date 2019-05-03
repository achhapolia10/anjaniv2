package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

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

//GetDeleteProducts Handler for route /delete requires id as a Params
func GetDeleteProducts(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	if id == "" {
		fmt.Print("No id Given")
		http.Redirect(w, req, "/products", 301)
	} else {
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Print("Id to Delete is not Proper")
			log.Fatal(err)
		} else {
			res := opdatabase.DeleteProduct(int64(i))
			if !res {
				fmt.Println("Error in Deleting Product")
			}
			http.Redirect(w, req, "/products", 301)
		}
	}
}
