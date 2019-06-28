package routes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/achhapolia10/anjaniv2/model"
	"github.com/achhapolia10/anjaniv2/opdatabase"
	"github.com/julienschmidt/httprouter"
)

//GetUsers Handler for /user
func GetUsers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) && isAdmin() {
		t, err := template.ParseGlob("views/components/navbar.comp")
		t, err = t.ParseFiles("views/users.html")
		if err != nil {
			log.Println(err)
		}
		users := model.GetUsers()
		data := struct {
			Users []opdatabase.User
			U     User
		}{
			users, currentUser,
		}
		err = t.ExecuteTemplate(w, "users.html", data)
		if err != nil {
			log.Println(err)
		}
	} else {
		http.Redirect(w, req, "/logout", http.StatusFound)
	}
}

//PostChange Handler for /user/change
func PostChange(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}

//PostDelete Handler for /user/change
func PostDelete(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}

//GetNewUser Handler for /user/change
func GetNewUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}
