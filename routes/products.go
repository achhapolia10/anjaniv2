//Package routes For Products
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

var groups map[int]string

//GetProducts Handler for route / method GET
func GetProducts(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {

		groups = make(map[int]string)
		gs, _ := model.GetGroups()
		for _, g := range gs {
			groups[g.Id] = g.Name
		}
		t := template.Must(template.ParseGlob("views/components/navbar.comp"))
		t.Funcs(template.FuncMap{
			"getGroupName": GetGroupName,
		}).ParseFiles("views/products.html")

		p, r := model.GetAllProduct()
		if r {
			d := struct {
				Products []opdatabase.Product
				U        User
			}{p, currentUser}
			err := t.ExecuteTemplate(w, "products.html", d)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

//GetDeleteProducts Handler for route /delete requires id as a Params
func GetDeleteProducts(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	if isLoggedIn(w, req) {

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
				res := opdatabase.DeleteProduct(i)
				if !res {
					fmt.Println("Error in Deleting Product")
				}
				http.Redirect(w, req, "/products", 301)
			}
		}
	}
}

//GetEditProduct handler /edit/:id
//id is a number for product id
func GetEditProduct(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	if isLoggedIn(w, req) {
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
				p, res := opdatabase.SelectProductID(i)
				if !res {
					fmt.Println("Error in Getting Product")
				} else {
					data := struct {
						Product opdatabase.Product
						U       User
					}{p, currentUser}
					t := template.Must(template.ParseGlob("views/components/navbar.comp"))
					t.ParseFiles("views/editProduct.html")
					t.Funcs(template.FuncMap{
						"getGroupName": GetGroupName,
					})
					t.ExecuteTemplate(w, "editProduct.html", data)

				}

			}
		}
	}
}

//PostEditProduct handler /edit/:id
//id is a number form produt id
func PostEditProduct(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	if isLoggedIn(w, req) {
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
				var pr, weight float64
				n = req.FormValue("product")
				p, err = strconv.Atoi(req.FormValue("box"))
				b, err = strconv.Atoi(req.FormValue("packet"))
				op, err = strconv.Atoi(req.FormValue("opacket"))
				ob, err = strconv.Atoi(req.FormValue("obox"))
				weight, err = strconv.ParseFloat(req.FormValue("weight"))
				pr, err = strconv.ParseFloat(req.FormValue("price"), 32)
				product := opdatabase.Product{
					ID: i, Name: n, PacketQuantity: p, BoxQuantity: b, Price: pr, OpeningBox: ob, OpeningPacket: op, Group: 0, Weight: weight,
				}
				result := opdatabase.EditProduct(i, product)
				if result {
					log.Println("Product Edited")
				} else {
					log.Println("Product Not Edited")
				}
				http.Redirect(w, req, "/products", 301)

			}
		}

	}
}

//GetNewProduct Handler for route / method GET
func GetNewProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {

		t := template.Must(template.ParseGlob("views/components/navbar.comp"))
		t.ParseFiles("views/newProduct.html")
		g, _ := model.GetGroups()
		data := struct {
			Groups []opdatabase.Group
			U      User
		}{
			g, currentUser,
		}
		t.ExecuteTemplate(w, "newProduct.html", data)
	}
}

//PostNewProduct creates a new Product
func PostNewProduct(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		err := req.ParseForm()
		if err != nil {
			log.Println(err)
		}
		var n string
		var p, b, op, ob, g int
		var pr, we float64
		n = req.FormValue("product")
		b, err = strconv.Atoi(req.FormValue("box"))
		p, err = strconv.Atoi(req.FormValue("packet"))
		op, err = strconv.Atoi(req.FormValue("opacket"))
		ob, err = strconv.Atoi(req.FormValue("obox"))
		pr, err = strconv.ParseFloat(req.FormValue("price"), 32)
		g, err = strconv.Atoi(req.FormValue("product-group"))
		we, err = strconv.ParseFloat(req.FormValue("weight"), 32)

		product := opdatabase.Product{
			ID: 0, Name: n, PacketQuantity: p,
			BoxQuantity: b, Price: pr, OpeningBox: ob, OpeningPacket: op, Group: g, Weight: we,
		}
		opdatabase.AddProduct(product)
		http.Redirect(w, req, "/products", 301)
	}
}

//GetGroupName returns name of the group
func GetGroupName(id int) string {
	return groups[id]
}
