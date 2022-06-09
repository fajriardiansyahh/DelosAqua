package main

import (
	"golang/controllers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.HomePage)
	mux.HandleFunc("/question", controllers.QuestionsPage)

	log.Println("Starting Web Server on port 9000")

	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}
