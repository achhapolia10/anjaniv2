package routes

import (
	"html/template"
	"net/http"

	"github.com/achhapolia10/anjaniv2/model"

	"github.com/julienschmidt/httprouter"
)

//GetDispatch route /dispatch method:GET
func GetDispatch(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	p, _ := model.GetAllProduct()
	t := template.Must(template.ParseGlob("views/components/*.comp"))
	t.ParseFiles("views/dispatch.html")
	t.ExecuteTemplate(w, "dispatch.html", p)
}

//PostDispatchNew route /dispacth/new method:Post
func PostDispatchNew(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

}
