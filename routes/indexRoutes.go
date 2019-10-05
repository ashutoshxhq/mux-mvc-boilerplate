package routes

import (
	"github.com/gorilla/mux"
	"webkodes.com/server/handlers"
)

// IndexRouter created auth router
func IndexRouter(r *mux.Router) *mux.Router {
	r.HandleFunc("/", handlers.IndexGetHandler).Methods("GET")
	r.HandleFunc("/about", handlers.AboutGetHandler).Methods("GET")
	return r
}
