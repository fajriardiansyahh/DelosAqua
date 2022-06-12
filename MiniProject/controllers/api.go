package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"miniproject/models"
	"net/http"
	"strconv"
)

func Farms_API(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
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

		w.Write(jsonResp)

	case http.MethodPost:
		jsonResp, err := json.Marshal(models.CreateFarm("API Test", "API Test Description", "API Test Thumbnails", w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}
		w.Write(jsonResp)

	case http.MethodPut:
		jsonResp, err := json.Marshal(models.UpdateFarm(4, "test create from update", "test create from update", "", w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}
		w.Write(jsonResp)

	case http.MethodDelete:
		jsonResp, err := json.Marshal(models.DeleteFarm(3, w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}
		w.Write(jsonResp)

	default:

	}
}

func Ponds_API(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
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
		w.Write(jsonResp)

	case http.MethodPost:
		jsonResp, err := json.Marshal(models.CreatePonds(1, "test", "test", "", w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}
		w.Write(jsonResp)

	case http.MethodPut:
		jsonResp, err := json.Marshal(models.UpdatePonds(4, 1, "test create from update", "test create from update", "", w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}
		w.Write(jsonResp)

	case http.MethodDelete:
		jsonResp, err := json.Marshal(models.DeletePonds(3, w))
		if err != nil {
			log.Println(err)
			w.Header().Set("Status", fmt.Sprint(http.StatusInternalServerError))
			http.Error(w, models.GetFarm_All(w).Message, models.GetFarm_All(w).Status)
			return
		}
		w.Write(jsonResp)

	default:

	}
}
