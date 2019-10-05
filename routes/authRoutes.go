package routes

import (
	"github.com/gorilla/mux"
	"webkodes.com/server/handlers"
)

// AuthRouter created auth router
func AuthRouter(r *mux.Router) *mux.Router {
	r.HandleFunc("/login", handlers.LoginGetHandler).Methods("GET")
	r.HandleFunc("/signup", handlers.SignupGetHandler).Methods("GET")
	return r
}
