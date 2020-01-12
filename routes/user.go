package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/achhapolia10/inventory-manager/model"
	"github.com/achhapolia10/inventory-manager/opdatabase"
	"github.com/julienschmidt/httprouter"
)

//GetUsers Handler for /user
func GetUsers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		if isAdmin() {
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
}

//GetChange Handler for /user/change
func GetChange(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		if isAdmin() {
			t, err := template.ParseGlob("views/components/navbar.comp")
			t, err = t.ParseFiles("views/cusers.html")
			if err != nil {
				log.Println(err)
			}
			username := req.URL.Query().Get("username")
			data := struct {
				Users string
				U     User
			}{
				username, currentUser,
			}
			err = t.ExecuteTemplate(w, "cusers.html", data)
			if err != nil {
				log.Println(err)
			}
		} else {
			http.Redirect(w, req, "/logout", http.StatusFound)
		}
	}
}

//PostChange Handler for /user/change
func PostChange(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		if isAdmin() {
			u := req.FormValue("username")
			p := req.FormValue("password")

			model.UpdatePassword(u, p)

		}
	}
}

//PostDelete Handler for /user/change
func PostDelete(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) {
		if isAdmin() {
			u := req.URL.Query().Get("username")

			if u != "admin" {
				model.DeleteUser(u)
			}

		}
		http.Redirect(w, req, "/users", http.StatusFound)
	}

}

//GetNewUser Handler for /user/change
func GetNewUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if isLoggedIn(w, req) && isAdmin() {
		t := template.Must(template.ParseGlob("views/components/navbar.comp"))
		t.ParseFiles("views/newusers.html")
		data := struct {
			U User
		}{
			currentUser,
		}
		t.ExecuteTemplate(w, "newusers.html", data)
	}
}

//PostNewUser Handler for /user/change
func PostNewUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := req.FormValue("username")
	p := req.FormValue("password")
	a := req.FormValue("user")

	admin, _ := strconv.Atoi(a)

	model.CreateUser(u, p, admin)

	http.Redirect(w, req, "/users", http.StatusFound)
}
