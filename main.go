package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"webkodes.com/server/routes"
	"webkodes.com/server/utils"
)

func main() {
	utils.LoadTemplates("views/*.html")
	r := mux.NewRouter()

	r = routes.IndexRouter(r)
	r = routes.AuthRouter(r)
	fs := http.FileServer(http.Dir("./public/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
