package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Farm struct {
	ID          int
	Name        string
	Description sql.NullString
	Thumbnails  sql.NullString
	Created_at  sql.NullString
	Updated_at  sql.NullString
	Deleted_at  sql.NullString
}

type Response_Farm struct {
	Status  int
	Message string
	Data    []Farm
}

func connection(w http.ResponseWriter) (*sql.DB, error) {

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

func CreateFarm(name string, description string, thumbnails string, w http.ResponseWriter) Response_Farm {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer connecting.Close()
	data, err := connecting.Query("INSERT INTO farm (name, description, thumbnails) VALUES (?,?,?)", name, description, thumbnails)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer data.Close()
	var result []Farm

	for data.Next() {
		var farm = Farm{}
		var err = data.Scan(&farm.ID, &farm.Name, &farm.Thumbnails, &farm.Description, &farm.Created_at, &farm.Updated_at, &farm.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Farm{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Farm{},
			}
		}
		result = append(result, farm)
	}

	return Response_Farm{
		Status:  http.StatusOK,
		Message: "CreateFarm",
		Data:    result,
	}
}

func UpdateFarm(id int, name string, description string, thumbnails string, w http.ResponseWriter) Response_Farm {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer connecting.Close()
	data_check, err := connecting.Query("SELECT * FROM farm WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return Response_Farm{
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    []Farm{},
			}
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer data_check.Close()

	if !data_check.Next() {
		return CreateFarm(name, description, thumbnails, w)
	}

	data, err := connecting.Query("UPDATE farm SET name=?, description=?, thumbnails=?, updated_at=?  WHERE id=? RETURN id, name, description, thumbnails", name, description, thumbnails, time.Now(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return Response_Farm{
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    []Farm{},
			}
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer data.Close()
	var result []Farm

	for data.Next() {
		var farm = Farm{}
		var err = data.Scan(&farm.ID, &farm.Name, &farm.Thumbnails, &farm.Description, &farm.Created_at, &farm.Updated_at, &farm.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Farm{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Farm{},
			}
		}
		result = append(result, farm)
	}

	return Response_Farm{
		Status:  http.StatusOK,
		Message: "UpdateFarm",
		Data:    result,
	}
}

func DeleteFarm(id int, w http.ResponseWriter) Response_Farm {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer connecting.Close()
	data_check, err := connecting.Query("SELECT * FROM farm WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return Response_Farm{
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    []Farm{},
			}
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer data_check.Close()

	data, err := connecting.Query("UPDATE farm SET deleted_at=? WHERE id=? RETURN id, name, description, thumbnails", time.Now(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return Response_Farm{
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    []Farm{},
			}
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer data.Close()
	var result []Farm

	for data.Next() {
		var farm = Farm{}
		var err = data.Scan(&farm.ID, &farm.Name, &farm.Thumbnails, &farm.Description, &farm.Created_at, &farm.Updated_at, &farm.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Farm{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Farm{},
			}
		}
		result = append(result, farm)
	}

	return Response_Farm{
		Status:  http.StatusOK,
		Message: "DeleteFarm",
		Data:    result,
	}
}

func GetFarm_All(w http.ResponseWriter) Response_Farm {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer connecting.Close()

	data, err := connecting.Query("SELECT * FROM `farm` WHERE `deleted_at` IS NULL")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer data.Close()

	var result []Farm
	for data.Next() {
		var farm = Farm{}
		var err = data.Scan(&farm.ID, &farm.Name, &farm.Thumbnails, &farm.Description, &farm.Created_at, &farm.Updated_at, &farm.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Farm{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Farm{},
			}
		}
		result = append(result, farm)
	}

	w.WriteHeader(http.StatusOK)
	return Response_Farm{
		Status:  http.StatusOK,
		Message: "GetFarm_All",
		Data:    result,
	}
}

func GetFarm_ID(w http.ResponseWriter, id int) Response_Farm {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer connecting.Close()
	data, err := connecting.Query("SELECT * FROM farm WHERE id=?", id)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Farm{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Farm{},
		}
	}
	defer data.Close()

	var result []Farm

	for data.Next() {
		var farm = Farm{}
		var err = data.Scan(&farm.ID, &farm.Name, &farm.Thumbnails, &farm.Description, &farm.Created_at, &farm.Updated_at, &farm.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Farm{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Farm{},
			}
		}
		result = append(result, farm)
	}

	return Response_Farm{
		Status:  http.StatusOK,
		Message: "GetFarm_ID",
		Data:    result,
	}
}
