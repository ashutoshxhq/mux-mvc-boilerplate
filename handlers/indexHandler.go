package handlers

import (
	"net/http"

	"webkodes.com/server/utils"
)

// IndexGetHandler handles login get req
func IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "index.html", nil)
}

// AboutGetHandler handles signup get req
func AboutGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "about.html", nil)
}
