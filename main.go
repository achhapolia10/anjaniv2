package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/achhapolia10/anjaniv2/routes"

	"github.com/julienschmidt/httprouter"
)

//Index Handler Method GET
func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(w, "Hello")
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/login", routes.GetLogin)

	fmt.Println("Starting Server on Port: 4001")
	log.Fatal(http.ListenAndServe(":4001", router))
}
