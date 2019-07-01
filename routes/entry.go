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

//JournalResponse struct for json Response
type JournalResponse struct {
	JournalEntries []opdatabase.JournalEntry `json:"entries"`
	Labours        map[string]bool           `json:"labours"`
	Box            int                       `json:"tbox"`
	Packet         int                       `json:"tpacket"`
}

var labours map[string]bool

//GetEntry Handler for route / method GET
func GetEntry(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	labours = make(map[string]bool)

	if isLoggedIn(w, req) {
		p, res := model.GetAllProduct()
		if !res {
			log.Println("Error in querying all products in Entries")
		} else {
			data := struct {
				Products []opdatabase.Product
				U        User
			}{
				p, currentUser,
			}
			t := template.Must(template.ParseGlob("views/components/navbar.comp"))
			t.ParseFiles("views/entry.html")
			t.ExecuteTemplate(w, "entry.html", data)
		}
	}
}

//PostEntryNew Handler for route /new method POST
func PostEntryNew(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		q := req.URL.Query()
		box, _ := strconv.Atoi(q["box"][0])
		packet, _ := strconv.Atoi(q["packet"][0])
		id, _ := strconv.Atoi(q["product"][0])
		je := opdatabase.JournalEntry{
			ID:        0,
			Labour:    q["labour"][0],
			Date:      q["date"][0],
			Box:       box,
			Packet:    packet,
			ProductID: id,
		}
		go model.UpdateLabourNames(je.Labour, je.Date, labours)
		model.CreateJournalEntry(je)
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
}

//GetJournalEntriesAll Gets Journal Entry for a Product on a particular Date and Send Response
//Route /getall method GET
func GetJournalEntriesAll(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		q := req.URL.Query()

		ch := make(chan map[string]bool)
		go model.GetLabourNames(ch)

		date := q["date"][0]
		productID, err := strconv.Atoi(q["id"][0])
		if err != nil {
			log.Println("Error in GETJournal Entries all")
			log.Println(err)
		}
		je, box, packet, res := model.GetAllJournalEntry(date, productID)

		labours = <-ch

		result := JournalResponse{
			JournalEntries: je,
			Box:            box,
			Packet:         packet,
			Labours:        labours,
		}
		if res {
			p, err := json.Marshal(result)
			if err != nil {
				log.Println("Error in GetJournalEntries all in Marshalling")
				log.Println(err)
			}
			io.WriteString(w, string(p))
		}
	}
}

//PostDeleteJournalEntry Deletes a Journal Entry
// route /entry/delete?id=?&productID=? method POST
func PostDeleteJournalEntry(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		q := req.URL.Query()
		id, _ := strconv.Atoi(q["id"][0])
		productID, _ := strconv.Atoi(q["productID"][0])
		model.DeleteJournalEntry(id, productID)
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
}
