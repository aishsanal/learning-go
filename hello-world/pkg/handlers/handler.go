package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aishsanal/learning-go/hello-world/pkg/config"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

var appConfig config.AppConfig

func renderTemplate(w http.ResponseWriter, t string) {
	var template *template.Template
	if appConfig.UseCache {
		template = appConfig.TemplateCache[t]
	} else {
		cache, _ := CreateTemplateCache()
		template = cache[t]
	}

	err := template.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func SetConfig(a config.AppConfig) {
	appConfig = a
}

type Repository struct {
	app config.AppConfig
}

var repository *Repository

func SetRepository(repo *Repository) {
	repository = repo
}

func CreateRepository(cnf config.AppConfig) *Repository {
	return &Repository{
		app: cnf,
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pagesPath, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, path := range pagesPath {
		name := filepath.Base(path)
		templ, err := template.New(name).ParseFiles(path)
		if err != nil {
			return cache, err
		}

		layoutPath, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layoutPath) > 0 {
			templ, err = templ.ParseGlob(layoutPath[0])
			if err != nil {
				return cache, err
			}
		}
		cache[name] = templ
	}
	return cache, nil
}

// The below functions were written first and the above ones are the refined versions
var tempCache = make(map[string]*template.Template)

func renderTemplateOld(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	_, inMap := tempCache[t]
	if !inMap {
		log.Println("Loading template from disk")
		tmpl, err = loadTemplateOld(t)
		if err != nil {
			log.Println("Error loading template from disk")
			return
		}
		tempCache[t] = tmpl
	} else {
		log.Println("Loading template from template cache")
		tmpl = tempCache[t]
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error rendering template ", t)
	}
}

func loadTemplateOld(t string) (*template.Template, error) {
	templatesToLoad := []string{
		fmt.Sprintf("../../templates/%s", t),
		"../../templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templatesToLoad...)
	return tmpl, err
}

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("../../templates/"+tmpl, "../../templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error while rendering template:", err)
// 	}
// }
