package routes

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
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
	//p, res := get all products and pass it to template
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
		id,
	}
	//TODO:Create a Journal Entry
	res := Response{
		301,
		Response{20, ", "},
	}
	p, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	io.WriteString(w, string(p))

}

//GetJournalEntriesAll Gets Journal Entry for a Product on a particular Date and Send Response
//Route /getall method GET
func GetJournalEntriesAll(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	q := req.URL.Query()
	date := q["date"][0]
	id, err := strconv.Atoi(q["id"][0])
	if err != nil {
		log.Println("Error in GETJournal Entries all")
		log.Println(err)
	}
	je, res := opdatabase.SelectJournalEntry(date, id)
	if res {
		p, err := json.Marshal(je)
		if err != nil {
			log.Println("Error in GetJournalEntries all in Marshalling")
			log.Println(err)
		}
		io.WriteString(w, string(p))
	}
}

//PostDeleteJournalEntry Deletes a Journal Entry
// route /entry/delete?id=?&productid=? method POST
func PostDeleteJournalEntry(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	q := req.URL.Query()
	id, _ := strconv.Atoi(q["id"][0])
	productid, _ := strconv.Atoi(q["productid"][0])
	opdatabase.DeleteJournalEntry(productid, id)
	res := Response{
		301,
		Response{20, ", "},
	}
	p, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	io.WriteString(w, string(p))
}
