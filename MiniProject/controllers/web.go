package controllers

import (
	"html/template"
	"log"
	"miniproject/models"
	"net/http"
	"path"
	"strconv"

	"github.com/fatih/structs" // An external package for handling structs data type into map[]bytes
)

// A function for handling Home view with URL representative "url.com"
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
	data := models.GetFarm_All(w)
	err_view := template.Execute(w, structs.Map(data))
	if err_view != nil {
		log.Println(err_view.Error())
		http.Error(w, "Error View", http.StatusInternalServerError)
		return
	}

	log.Println(template.Tree.Name)
}

// A function for handling Farms Model details view
func Farms_view(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	id := r.URL.Query().Get("id")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	val_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	template, err_func := template.ParseFiles(path.Join("views/models/farm", "view.html"), path.Join("views/models", "index.html"), path.Join("views/layouts", "apps.html"))
	if err_func != nil {
		log.Println(err_func.Error())
		http.Error(w, err_func.Error(), http.StatusInternalServerError)
		return
	}
	data := models.GetFarm_ID(w, val_id)
	if len(data.Data) < 1 {
		log.Println(http.StatusNotFound)
		http.NotFound(w, r)
		return
	}

	err_view := template.Execute(w, data.Data)
	if err_view != nil {
		log.Println(err_view.Error())
		http.Error(w, err_view.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(template.Tree.Name)
}

// A function for handling Ponds Model details view
func Ponds_view(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
}
