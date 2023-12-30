package main

import (
	"fmt"
	"log"
	"net/http"
	"nylespham/airbnb/packages/config"
	"nylespham/airbnb/packages/handlers"
	"nylespham/airbnb/packages/render"
)

const port = ":8080"

func main() {
	var app config.AppConfig
	templateCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Server is running on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
