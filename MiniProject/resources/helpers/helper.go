package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql" // An external package for handling mysql
)

// Connection function to connect the application into the Database
func Connection(w http.ResponseWriter) (*sql.DB, error) {

	connection, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/delosaqua")
	if err != nil {
		log.Println(err)
		jsonResp, err_rsp := json.Marshal(err)
		if err_rsp != nil {
			log.Println(err_rsp)
			w.WriteHeader(http.StatusInternalServerError)
			return nil, err_rsp
		}
		w.Write(jsonResp)
		return nil, err
	}

	connection.SetConnMaxLifetime(time.Minute * 5)
	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(10)

	return connection, nil
}

func API_Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path + r.URL.RawQuery)
	w.Header().Set("Content-Type", "application/json")
	accept := r.Header.Get("Accept")
	contentType := r.Header.Get("Content-Type")
	if accept != "application/json" || contentType != "application/json" {
		msg := "Content-Type / Accept header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	r.ParseForm()
}

func API_Responses(data any, w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
		return
	}

	w.Write(jsonResp)
}
