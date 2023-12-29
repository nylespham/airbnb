package handlers

import (
	"net/http"
	"nylespham/airbnb/packages/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "/home")
}

// About is a function that returns the about page
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "/about")
}
