package models

import (
	"database/sql"
	"log"
	"miniproject/resources/helpers"
	"net/http"
	"time"
)

// Struct as a container to the representative table on the DB
type Farm struct {
	ID          int
	Name        string
	Description string
	Thumbnails  string
	Created_at  sql.NullString
	Updated_at  sql.NullString
	Deleted_at  sql.NullString
}

// Struct as a store to contain the response from function on Farms Models
type Response_Farm struct {
	// Status is a representation to determine if the function succeeded or not
	Status int
	// Message to serve the explanation about what is going on
	Message string
	Data    []Farm
}

// Function to store Farms data into Database
func CreateFarm(name string, description string, thumbnails string, w http.ResponseWriter) Response_Farm {
	connecting, err := helpers.Connection(w)
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

// Function to updating Farms data into Database
func UpdateFarm(id int, name string, description string, thumbnails string, w http.ResponseWriter) Response_Farm {
	connecting, err := helpers.Connection(w)
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

	data, err := connecting.Query("UPDATE farm SET name=?, description=?, thumbnails=?, updated_at=?  WHERE id=?", name, description, thumbnails, time.Now(), id)
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

// Function to soft deleting Farms data into Database
func DeleteFarm(id int, w http.ResponseWriter) Response_Farm {
	connecting, err := helpers.Connection(w)
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

	data, err := connecting.Query("UPDATE farm SET deleted_at=? WHERE id=?", time.Now(), id)
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

// Function to get all Farms record data from Database
func GetFarm_All(w http.ResponseWriter) Response_Farm {
	connecting, err := helpers.Connection(w)
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

// Function to get Farms record data from Database by their ID
func GetFarm_ID(w http.ResponseWriter, id int) Response_Farm {
	connecting, err := helpers.Connection(w)
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
