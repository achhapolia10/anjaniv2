package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/achhapolia10/anjaniv2/model"

	"github.com/julienschmidt/httprouter"
)

//GetProducts Handler for route / method GET
func GetProducts(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/products.html")
	//TODO : p, r := opdatabase.SelectProduct()
	if r {
		t.ExecuteTemplate(w, "products.html", p)
	}
}

//GetDeleteProducts Handler for route /delete requires id as a Params
func GetDeleteProducts(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	id := p.ByName("id")
	if id == "" {
		fmt.Println("No id Given")
		http.Redirect(w, req, "/products", 301)
	} else {
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Id to Delete is not Proper")
			log.Println(err)
		} else {
			//TODO res := opdatabase.DeleteProduct(int64(i))
			if !res {
				fmt.Println("Error in Deleting Product")
			}
			http.Redirect(w, req, "/products", 301)
		}
	}
}

//GetEditProduct handler /edit/:id
//id is a number for product id
func GetEditProduct(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if id == "" {
		fmt.Println("No id Given")
		http.Redirect(w, req, "/products", 301)
	} else {
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Id to Edit is not Proper")
			log.Println(err)
		} else {
			//TODO : p, res := opdatabase.SelectProductID(i)
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
		fmt.Println("No id Given")
		http.Redirect(w, req, "/products", 301)
	} else {
		i, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Id to Edit is not Proper")
			log.Println(err)
		} else {
			err = req.ParseForm()
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
				i, n, p, b, pr, ob, op,
			}
			//TODO : result := opdatabase.EditProduct(i, product)
			if result {
				log.Println("Product Edited")
			} else {
				log.Println("Product Not Edited")
			}
			http.Redirect(w, req, "/products", 301)

		}
	}

}

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
	//TODO : opdatabase.AddProduct(product)
	http.Redirect(w, req, "/products", 301)
}
