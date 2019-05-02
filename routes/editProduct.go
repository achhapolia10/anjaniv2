package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/achhapolia10/anjaniv2/model"
	"github.com/achhapolia10/anjaniv2/opdatabase"
	"github.com/julienschmidt/httprouter"
)

//GetEditProduct handler /edit/:id
//id is a number for product id
func GetEditProduct(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if id == "" {
		fmt.Print("No id Given")
		http.Redirect(w, req, "/products", 301)
	} else {
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Print("Id to Edit is not Proper")
			log.Fatal(err)
		} else {
			p, res := opdatabase.SelectProductID(i)
			if !res {
				fmt.Println("Error in Getting Product")
			} else {
				t := template.Must(template.ParseGlob("views/components/*.comp"))
				t.ParseFiles("views/editProduct.html")
				t.ExecuteTemplate(w, "editProduct.html", p)

			}

		}
	}
}

//PostEditProduct handler /edit/:id
//id is a number form produt id
func PostEditProduct(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if id == "" {
		fmt.Print("No id Given")
		http.Redirect(w, req, "/products", 301)
	} else {
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Print("Id to Edit is not Proper")
			log.Fatal(err)
		} else {
			err = req.ParseForm()
			if err != nil {
				log.Fatal(err)
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
				i, n, p, b, pr, ob, op,
			}
			result := opdatabase.EditProduct(i, product)
			if result {
				fmt.Print("Product Edited")
			} else {
				fmt.Print("Product Not Edited")
			}
			http.Redirect(w, req, "/products", 301)

		}
	}
}
