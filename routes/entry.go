package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/achhapolia10/anjaniv2/opdatabase"

	"github.com/julienschmidt/httprouter"
)

//Response struct for json Responses
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

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

//PostEntryNew Handler for route /new method POST
func PostEntryNew(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	a := req.URL.Query()
	fmt.Print(a)

	res := Response{
		301,
		Response{20, ", "},
	}
	p, err := json.Marshal(res)
	if err != nil {
		log.Print(err)
	}
	io.WriteString(w, string(p))
	os.Stdout.Write(p)

}
