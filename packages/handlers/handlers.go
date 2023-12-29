package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/home.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template: ", err)
		return
	}
}

// About is a function that returns the about page
func About(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/about.html")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error executing template: ", err)
		return
	}
}
