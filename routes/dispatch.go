package routes

import (
	"encoding/json"
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
	p, _ := model.GetAllProduct()
	initializeProductsMap(p)
	t := template.Must(template.ParseGlob("views/components/navbar.comp"))
	t.ParseFiles("views/dispatch.html")
	t.ExecuteTemplate(w, "dispatch.html", p)
}

//PostDispatchNew route /dispacth/new method:Post
func PostDispatchNew(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	req.ParseForm()
	productID, _ := strconv.Atoi(req.FormValue("product"))
	box, _ := strconv.Atoi(req.FormValue("box"))
	packet, _ := strconv.Atoi(req.FormValue("packet"))
	date := req.FormValue("datev")
	se := opdatabase.StockEntry{
		0, date, 0, 0, box, packet, productID,
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

//GetDispatchEntries returns the stock entries for date
func GetDispatchEntries(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	value := req.URL.Query()
	date := value.Get("date")
	s := model.GetDispatchEntriesByDate(date)
	t := template.New("dispatch")
	t.Funcs(template.FuncMap{
		"getProductName": GetProductName,
	}).ParseFiles("views/components/dispatch.comp")
	t.ExecuteTemplate(w, "dispatch.comp", s)
}

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
