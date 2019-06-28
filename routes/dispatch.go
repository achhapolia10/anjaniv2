package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/achhapolia10/anjaniv2/opdatabase"

	"github.com/achhapolia10/anjaniv2/model"

	"github.com/julienschmidt/httprouter"
)

var products map[int]string

//GetDispatch route /dispatch method:GET
func GetDispatch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		p, _ := model.GetAllProduct()
		initializeProductsMap(p)
		t := template.Must(template.ParseGlob("views/components/navbar.comp"))
		t.ParseFiles("views/dispatch.html")
		data := struct {
			Products []opdatabase.Product
			U        User
		}{
			p, currentUser,
		}
		t.ExecuteTemplate(w, "dispatch.html", data)
		return
	}
}

//PostDispatchNew route /dispacth/new method:Post
func PostDispatchNew(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		req.ParseForm()
		productID, _ := strconv.Atoi(req.FormValue("product"))
		box, _ := strconv.Atoi(req.FormValue("box"))
		packet, _ := strconv.Atoi(req.FormValue("packet"))
		date := req.FormValue("date")
		se := opdatabase.StockEntry{
			ID: 0, Date: date, BoxIn: 0, PacketIn: 0, BoxOut: box, PacketOut: packet, ProductID: productID,
		}
		res := model.NewDispatchEntry(se)
		if res {
			r := Response{301, ""}
			p, err := json.Marshal(r)
			if err != nil {
				log.Println("Error in PostDispatchEntries all in Marshalling")
				log.Println(err)
			}
			io.WriteString(w, string(p))
		}
	}
}

//GetDispatchEntries returns the stock entries for date
func GetDispatchEntries(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		value := req.URL.Query()
		date := value.Get("date")
		s := model.GetDispatchEntriesByDate(date)
		t := template.New("dispatch")
		t.Funcs(template.FuncMap{
			"getProductName": GetProductName,
			"shouldPrint":    ShouldPrintD,
		}).ParseFiles("views/components/dispatch.comp")
		t.ExecuteTemplate(w, "dispatch.comp", s)
	}
}

//GetDispatchDelete delets an entry for date and product
func GetDispatchDelete(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		date := req.URL.Query().Get("date")
		productID, _ := strconv.Atoi(req.URL.Query().Get("product"))
		res := model.DeleteDispatchEntry(date, productID)
		if res {
			r := Response{301, ""}
			p, err := json.Marshal(r)
			if err != nil {
				fmt.Println("Errorin Marshalling response in GetDispatchDelete: ", err)
			}
			io.WriteString(w, string(p))
		}
	}
}

/*
 *Functions for Function Maps used in templates
 *
 */

func initializeProductsMap(p []opdatabase.Product) {
	products = make(map[int]string)
	for _, product := range p {
		products[product.ID] = product.Name
	}
}

//GetProductName returns name of Product for an id
func GetProductName(id int) string {
	return products[id]
}

//ShouldPrintD returns name true if a product need to be print
func ShouldPrintD(bo int, po int) bool {
	if bo == 0 && po == 0 {
		return false
	}
	return true
}
