package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"miniproject/models"
	"miniproject/resources/helpers"
	"net/http"
	"strconv"
)

type JSONBody struct {
	Data map[string]string
}

// A function to handle API method such as: GET, POST, PUT, DELETE for Farm Model
func Farms_API(w http.ResponseWriter, r *http.Request) {
	helpers.API_Handler(w, r)
	models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)

	switch r.Method {

	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id != "" {
			farm_id, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
				w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
				return
			}

			result := models.GetFarm_ID(w, farm_id)
			helpers.API_Responses(result, w, r)
			return
		}

		result := models.GetFarm_All(w)
		helpers.API_Responses(result, w, r)

	case http.MethodPost:
		body := models.Farm{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		result := models.CreateFarm(body.Name, body.Description, body.Thumbnails, w)
		helpers.API_Responses(result, w, r)

	case http.MethodPut:
		body := models.Farm{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		result := models.UpdateFarm(body.ID, body.Name, body.Description, body.Thumbnails, w)
		helpers.API_Responses(result, w, r)

	case http.MethodDelete:
		body := models.Farm{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		result := models.DeleteFarm(body.ID, w)
		helpers.API_Responses(result, w, r)

	default:

	}
}

// A function to handle API method such as: GET, POST, PUT, DELETE for Ponds Model
func Ponds_API(w http.ResponseWriter, r *http.Request) {
	helpers.API_Handler(w, r)
	models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)

	switch r.Method {

	case http.MethodGet:
		id := r.URL.Query().Get("id")
		if id != "" {
			ponds_id, err := strconv.Atoi(id)
			if err != nil {
				log.Println(err)
				w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
				return
			}

			result := models.GetPonds_ID(w, ponds_id)
			helpers.API_Responses(result, w, r)
			return
		}

		result := models.GetPonds_All(w)
		helpers.API_Responses(result, w, r)

	case http.MethodPost:
		body := models.Ponds{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		result := models.CreatePonds(
			body.Farm_ID,
			body.Name,
			body.Description,
			body.Thumbnails, w)
		helpers.API_Responses(result, w, r)

	case http.MethodPut:
		body := models.Ponds{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		result := models.UpdatePonds(
			body.ID,
			body.Farm_ID,
			body.Name,
			body.Description,
			body.Thumbnails, w)
		helpers.API_Responses(result, w, r)

	case http.MethodDelete:
		body := models.Ponds{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		result := models.DeletePonds(body.ID, w)
		helpers.API_Responses(result, w, r)

	default:

	}
}

func APIAnalyst_API(w http.ResponseWriter, r *http.Request) {
	helpers.API_Handler(w, r)
	models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)

	switch r.Method {
	case http.MethodGet:
		result := models.GetAPI_All(w)
		helpers.API_Responses(result, w, r)
	default:
		http.NotFound(w, r)
	}
}
