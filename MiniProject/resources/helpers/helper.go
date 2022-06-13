package helpers

import (
	"database/sql"
	"encoding/json"
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
