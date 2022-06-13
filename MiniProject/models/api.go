package models

import (
	"log"
	"miniproject/resources/helpers"
	"net/http"
)

// Struct as a container to the representative table on the DB
type API struct {
	ID         int
	Method     string
	Host       string
	Path       string
	UA         string
	Created_at string
}

// Struct as a store to contain the response from function on API Models
type Response_API struct {
	// Status is a representation to determine if the function succeeded or not
	Status int
	// Message to serve the explanation about what is going on
	Message string
	Data    []API
}

// Function to store API data into Database
func Create_API(method string, path string, host string, ua string, w http.ResponseWriter) Response_API {
	connecting, err := helpers.Connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_API{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []API{},
		}
	}
	defer connecting.Close()
	data, err := connecting.Query("INSERT INTO api (method, path, host, ua) VALUES (?,?,?,?)", method, path, host, ua)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_API{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []API{},
		}
	}
	defer data.Close()
	var result []API

	for data.Next() {
		var api = API{}
		var err = data.Scan(&api.ID, &api.Method, &api.Host, &api.Path, &api.UA, &api.Created_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_API{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []API{},
			}
		}
		result = append(result, api)
	}

	return Response_API{
		Status:  http.StatusOK,
		Message: "CreateAPI",
		Data:    result,
	}
}

// Function to remove API data from Database by their ID
func HardDelete_API(id int) Response_API {
	return Response_API{
		Status:  http.StatusOK,
		Message: "CreateAPI",
		Data:    []API{},
	}
}

// Function to get all API record data from Database
func GetAPI_All(w http.ResponseWriter) Response_API {
	connecting, err := helpers.Connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_API{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []API{},
		}
	}
	defer connecting.Close()

	data, err := connecting.Query("SELECT * FROM `api`")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_API{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []API{},
		}
	}
	defer data.Close()

	var result []API
	for data.Next() {
		var api = API{}
		var err = data.Scan(&api.ID, &api.Method, &api.Host, &api.Path, &api.UA, &api.Created_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_API{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []API{},
			}
		}
		result = append(result, api)
	}

	w.WriteHeader(http.StatusOK)
	return Response_API{
		Status:  http.StatusOK,
		Message: "GetAPI_All",
		Data:    result,
	}
}

func GetAPI_RefColumn() Response_API {
	return Response_API{
		Status:  http.StatusOK,
		Message: "CreateAPI",
		Data:    []API{},
	}
}

func GetAPI_Summary() Response_API {
	return Response_API{
		Status:  http.StatusOK,
		Message: "CreateAPI",
		Data:    []API{},
	}
}
