package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	const data = "nylespham"
	filePrefix, _ := filepath.Abs("../../templates/")
	parsedTemplate, err := template.ParseFiles(filePrefix+tmpl, filePrefix+"/base.layout.tmpl")
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
	parsedTemplate.Execute(w, data)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err := templateSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
			myCache[name] = ts
		}

	}
	return myCache, nil
}
