package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/achhapolia10/anjaniv2/opdatabase"

	"github.com/achhapolia10/anjaniv2/model"
	"github.com/julienschmidt/httprouter"
)

//GetGroup Handler for route / method: GET
func GetGroup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseGlob("views/components/navbar.comp"))
	t.ParseFiles("views/group.html")
	g, _ := model.GetGroups()
	log.Print(g)
	t.ExecuteTemplate(w, "group.html", g)
}

//PostGroupNew Handler for roue /new method :POST
func PostGroupNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error in Parsing form for a new Group: %v", err)
	}
	gn := r.FormValue("group-name-form")
	if gn != "" {
		g := opdatabase.Group{
			Id: 0, Name: gn,
		}
		model.NewGroup(g)
	} else {
		log.Printf("No Group Name Provided")
	}
	http.Redirect(w, r, "/groups", 301)
}

//PostGroupDelete Handler for roue /new method :POST
func PostGroupDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, _ := strconv.Atoi(params.ByName("id"))
	g := opdatabase.Group{
		Id: id, Name: "er",
	}
	model.DeleteGroup(g)
	http.Redirect(w, r, "/groups", 301)
}

//PostGroupEdit Handler for roue /new method :POST
func PostGroupEdit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := r.ParseForm()
	id, _ := strconv.Atoi(params.ByName("id"))
	if err != nil {
		log.Printf("Error in Parsing form for a new Group: %v", err)
	}
	gn := r.FormValue("group-name-form")
	if gn != "" {
		g := opdatabase.Group{
			Id: id, Name: gn,
		}
		model.EditGroup(g)
	} else {
		log.Printf("No Group Name Provided")
	}
	http.Redirect(w, r, "/groups", 301)
}
