package main

import (
	"io/ioutil"
	"log"
	"miniproject/controllers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/Farms", controllers.Farms_view)
	mux.HandleFunc("/Ponds", controllers.Ponds_view)

	mux.HandleFunc("/resources/js/apps.js", sendLocalJS)
	mux.HandleFunc("/resources/css/style.css", sendLocalCSS)

	mux.HandleFunc("/api/v1/farms", controllers.Farms_API)
	mux.HandleFunc("/api/v1/ponds", controllers.Ponds_API)

	log.Println("Starting Web Server at port: 9090")

	err := http.ListenAndServe(":9090", mux)
	log.Fatal(err.Error())
}

func sendLocalJS(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("resources/js/apps.js")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/javascript; charset=utf-8")
	w.Write(data)
}

func sendLocalCSS(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("resources/css/style.css")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Write(data)
}
