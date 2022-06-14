package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // An external package for handling mysql
	"github.com/joho/godotenv"         // An external package for get .env file variables
)

// Connection function to connect the application into the Database
func Connection(w http.ResponseWriter) (*sql.DB, error) {

	connection, err := sql.Open(load_Env(".env", "DATABASE_DRIVER"), load_Env(".env", "DATABASE_USERNAME")+":@"+load_Env(".env", "DATABASE_CONNECTION")+"("+load_Env(".env", "DATABASE_URL")+load_Env(".env", "DATABASE_PORT")+")/"+load_Env(".env", "DATABASE_NAME"))
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

// A helper to setup API Call
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

// A helper to set an api responses
func API_Responses(data any, w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
		return
	}

	w.Write(jsonResp)
}

// a local function as a helper to loading .env File with .env file name and key as a parameter and returning value of the key
func load_Env(file string, key string) string {
	err := godotenv.Load(file)
	if err != nil {
		panic(err)
	}
	val := os.Getenv(key)

	return val
}
