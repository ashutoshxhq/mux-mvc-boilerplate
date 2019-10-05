package handlers

import (
	"net/http"

	"webkodes.com/server/utils"
)

// LoginGetHandler handles login get req
func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "login.html", nil)
}

// SignupGetHandler handles signup get req
func SignupGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "signup.html", nil)
}
