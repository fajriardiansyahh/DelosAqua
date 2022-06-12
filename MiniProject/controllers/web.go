package controllers

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	template, err_func := template.ParseFiles(path.Join("views", "index.html"), path.Join("views/layouts", "apps.html"))
	if err_func != nil {
		log.Println(err_func.Error())
		http.Error(w, err_func.Error(), http.StatusInternalServerError)
		return
	}

	err_view := template.Execute(w, nil)
	if err_view != nil {
		log.Println(err_view.Error())
		http.Error(w, "Error View", http.StatusInternalServerError)
		return
	}

	log.Printf(template.Tree.Name)
}

func Farms_view(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
}

func Ponds_view(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

}
