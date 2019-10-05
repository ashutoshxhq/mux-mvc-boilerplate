package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"webkodes.com/server/routes"
	"webkodes.com/server/utils"
)

func main() {
	utils.LoadTemplates("views/*.html")
	r := mux.NewRouter()

	r = routes.IndexRouter(r)
	r = routes.AuthRouter(r)

	c := utils.GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	fs := http.FileServer(http.Dir("./public/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
