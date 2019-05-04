package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/achhapolia10/anjaniv2/model"

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
		log.Println("Error in querying all products in Entries")
	} else {
		t := template.Must(template.ParseGlob("views/components/*.comp"))
		t.ParseFiles("views/entry.html")
		t.ExecuteTemplate(w, "entry.html", p)
	}
}

//PostEntryNew Handler for route /new method POST
func PostEntryNew(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	q := req.URL.Query()
	box, _ := strconv.Atoi(q["box"][0])
	packet, _ := strconv.Atoi(q["packet"][0])
	id, _ := strconv.Atoi(q["product"][0])
	je := model.JournalEntry{
		0,
		model.Labour(q["labour"][0]),
		q["date"][0],
		box,
		packet,
		&model.Product{
			id, "", 0, 0, 0, 0, 0,
		},
	}
	opdatabase.NewJournalEntry(je)
	res := Response{
		301,
		Response{20, ", "},
	}
	p, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	io.WriteString(w, string(p))
	os.Stdout.Write(p)

}
