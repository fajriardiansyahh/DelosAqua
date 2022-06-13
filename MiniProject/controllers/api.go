package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"miniproject/models"
	"net/http"
	"strconv"
)

type JSONBody struct {
	Data map[string]string
}

// A function to handle API method such as: GET, POST, PUT, DELETE for Farm Model
func Farms_API(w http.ResponseWriter, r *http.Request) {
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

			jsonResp, err := json.Marshal(models.GetFarm_ID(w, farm_id))
			if err != nil {
				log.Println(err)
				w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
				http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
				return
			}
			models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
			w.Write(jsonResp)
			return
		}

		jsonResp, err := json.Marshal(models.GetFarm_All(w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}

		models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
		w.Write(jsonResp)

	case http.MethodPost:
		body := models.Farm{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		result := models.CreateFarm(body.Name, body.Description, body.Thumbnails, w)
		jsonResp, err := json.Marshal(result)
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			w.Write([]byte(err.Error()))
			return
		}

		models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
		w.Write(jsonResp)

	case http.MethodPut:
		body := models.Farm{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}
		result := models.UpdateFarm(body.ID, body.Name, body.Description, body.Thumbnails, w)
		jsonResp, err := json.Marshal(result)
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}

		models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
		w.Write(jsonResp)

	case http.MethodDelete:
		body := models.Farm{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		jsonResp, err := json.Marshal(models.DeleteFarm(body.ID, w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}

		models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
		w.Write(jsonResp)

	default:

	}
}

// A function to handle API method such as: GET, POST, PUT, DELETE for Ponds Model
func Ponds_API(w http.ResponseWriter, r *http.Request) {
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

			jsonResp, err := json.Marshal(models.GetPonds_ID(w, ponds_id))
			if err != nil {
				log.Println(err)
				w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
				return
			}

			models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
			w.Write(jsonResp)
			return
		}

		jsonResp, err := json.Marshal(models.GetPonds_All(w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}

		models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
		w.Write(jsonResp)

	case http.MethodPost:
		body := models.Ponds{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		jsonResp, err := json.Marshal(models.CreatePonds(body.Farm_ID, body.Name, body.Description, body.Thumbnails, w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}

		models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
		w.Write(jsonResp)

	case http.MethodPut:
		body := models.Ponds{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		jsonResp, err := json.Marshal(models.UpdatePonds(body.ID, body.Farm_ID, body.Name, body.Description, body.Thumbnails, w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}

		models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
		w.Write(jsonResp)

	case http.MethodDelete:
		body := models.Ponds{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err.Error())
		}

		jsonResp, err := json.Marshal(models.DeletePonds(body.ID, w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}

		models.Create_API(r.Method, r.URL.Path+r.URL.RawQuery, r.Host, r.UserAgent(), w)
		w.Write(jsonResp)

	default:

	}
}
