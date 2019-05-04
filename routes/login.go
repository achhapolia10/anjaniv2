package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//GetLogin Handler for route /login method Get
func GetLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("views/login.html")
	if err != nil {
		log.Printlnln(err)
	}
	t.Execute(w, "")
}

//PostLogin Handler for route /login method Post
//Will Handle the login
func PostLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	//TODO: LoginLogic
	http.Redirect(w, req, "/", 301)
}
