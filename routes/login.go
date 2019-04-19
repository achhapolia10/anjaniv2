package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//GetLogin Handler for route /login method Get
func GetLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("views/login.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, "")
}
