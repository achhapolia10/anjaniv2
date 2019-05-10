package routes

import (
	"html/template"
	"net/http"

	"github.com/achhapolia10/anjaniv2/model"
	"github.com/julienschmidt/httprouter"
)

//GetGroup Handler for route / method: GET
func GetGroup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/group.html")
	g, _ := model.GetGroups()
	t.ExecuteTemplate(w, "group.html", g)
}

//GetGroupNew Handler for route /new method :GET
func GetGroupNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/newgroup.html")
	t.ExecuteTemplate(w, "groups.html", "")
}

//PostGroupNew Handler for roue /new method :POST
func PostGroupNew(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

//PostGroupDelete Handler for roue /new method :POST
func PostGroupDelete(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

//GetGroupEdit Handler for roue /new method :POST
func GetGroupEdit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

//PostGroupEdit Handler for roue /new method :POST
func PostGroupEdit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
