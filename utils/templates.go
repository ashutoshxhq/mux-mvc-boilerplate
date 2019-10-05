package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates exported load template
func LoadTemplates(pattern string) {
	templates = template.Must(template.ParseGlob(pattern))
}

// ExecuteTemplate exported execute template
func ExecuteTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

// Render exported render template
func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}
