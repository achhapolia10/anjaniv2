package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/achhapolia10/anjaniv2/model"

	"github.com/julienschmidt/httprouter"
)

//User Defines the User Type
type User struct {
	Username   string
	Admin      int
	IsLoggedIn bool
}

var currentUser User

//GetLogin Handler for route /login method Get
func GetLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Print(currentUser)
	if !currentUser.IsLoggedIn {
		t, err := template.ParseFiles("views/login.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, "")
	} else {
		http.Redirect(w, req, "/stock", http.StatusFound)
	}
}

//PostLogin Handler for route /login method Post
//Will Handle the login
func PostLogin(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if !currentUser.IsLoggedIn {
		uname := req.FormValue("username")
		password := req.FormValue("password")
		admin, r := model.CheckUser(uname, password)
		if r {
			currentUser = User{
				Username:   uname,
				Admin:      admin,
				IsLoggedIn: true,
			}
			http.Redirect(w, req, "/stock", http.StatusFound)
			return
		} else {
			currentUser.IsLoggedIn = false
		}
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

//GetLogout Handler for route /logout method Get
func GetLogout(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	currentUser.IsLoggedIn = false
	http.Redirect(w, req, "/", http.StatusFound)
}

func isLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	if currentUser.IsLoggedIn {
		return true
	}
	http.Redirect(w, req, "/", http.StatusFound)
	return false
}

func isAdmin() bool {
	if currentUser.Admin == 1 {
		return true
	}
	return false
}
