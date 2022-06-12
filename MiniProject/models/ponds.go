package models

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

type Ponds struct {
	ID          int
	Farm_ID     int
	Name        string
	Description sql.NullString
	Thumbnails  sql.NullString
	Created_at  sql.NullString
	Updated_at  sql.NullString
	Deleted_at  sql.NullString
}

type Response_Ponds struct {
	Status  int
	Message string
	Data    []Ponds
}

func CreatePonds(farm_id int, name string, description string, thumbnails string, w http.ResponseWriter) Response_Ponds {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer connecting.Close()
	data, err := connecting.Query("INSERT INTO ponds (farm_id, name, description, thumbnails) VALUES (?,?,?,?)", farm_id, name, description, thumbnails)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer data.Close()
	var result []Ponds

	for data.Next() {
		var ponds = Ponds{}
		var err = data.Scan(&ponds.ID, &ponds.Name, &ponds.Thumbnails, &ponds.Description, &ponds.Created_at, &ponds.Updated_at, &ponds.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Ponds{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		result = append(result, ponds)
	}

	return Response_Ponds{
		Status:  http.StatusOK,
		Message: "CreatePonds",
		Data:    result,
	}
}

func UpdatePonds(id int, farm_id int, name string, description string, thumbnails string, w http.ResponseWriter) Response_Ponds {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer connecting.Close()
	data_check, err := connecting.Query("SELECT * FROM ponds WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return Response_Ponds{
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer data_check.Close()

	if !data_check.Next() {
		return CreatePonds(farm_id, name, description, thumbnails, w)
	}

	data, err := connecting.Query("UPDATE ponds SET farm_id=?, name=?, description=?, thumbnails=?, updated_at=?  WHERE id=?", farm_id, name, description, thumbnails, time.Now(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return Response_Ponds{
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer data.Close()
	var result []Ponds

	for data.Next() {
		var ponds = Ponds{}
		var err = data.Scan(&ponds.ID, &ponds.Name, &ponds.Thumbnails, &ponds.Description, &ponds.Created_at, &ponds.Updated_at, &ponds.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Ponds{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		result = append(result, ponds)
	}

	return Response_Ponds{
		Status:  http.StatusOK,
		Message: "UpdatePonds",
		Data:    result,
	}
}

func DeletePonds(id int, w http.ResponseWriter) Response_Ponds {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer connecting.Close()
	data_check, err := connecting.Query("SELECT * FROM ponds WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return Response_Ponds{
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer data_check.Close()

	data, err := connecting.Query("UPDATE ponds SET deleted_at=? WHERE id=?", time.Now(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return Response_Ponds{
				Status:  http.StatusNotFound,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer data.Close()
	var result []Ponds

	for data.Next() {
		var ponds = Ponds{}
		var err = data.Scan(&ponds.ID, &ponds.Name, &ponds.Thumbnails, &ponds.Description, &ponds.Created_at, &ponds.Updated_at, &ponds.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Ponds{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		result = append(result, ponds)
	}

	return Response_Ponds{
		Status:  http.StatusOK,
		Message: "DeletePonds",
		Data:    result,
	}
}

func GetPonds_All(w http.ResponseWriter) Response_Ponds {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer connecting.Close()
	data, err := connecting.Query("SELECT * FROM ponds WHERE deleted_at IS NULL")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer data.Close()
	var result []Ponds

	for data.Next() {
		var ponds = Ponds{}
		var err = data.Scan(&ponds.ID, &ponds.Farm_ID, &ponds.Name, &ponds.Thumbnails, &ponds.Description, &ponds.Created_at, &ponds.Updated_at, &ponds.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Ponds{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		result = append(result, ponds)
	}

	return Response_Ponds{
		Status:  http.StatusOK,
		Message: "GetPonds_All",
		Data:    result,
	}
}

func GetPonds_ID(w http.ResponseWriter, id int) Response_Ponds {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer connecting.Close()
	data, err := connecting.Query("SELECT * FROM ponds WHERE id=?", id)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer data.Close()

	var result []Ponds

	for data.Next() {
		var ponds = Ponds{}
		var err = data.Scan(&ponds.ID, &ponds.Farm_ID, &ponds.Name, &ponds.Thumbnails, &ponds.Description, &ponds.Created_at, &ponds.Updated_at, &ponds.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Ponds{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		result = append(result, ponds)
	}

	return Response_Ponds{
		Status:  http.StatusOK,
		Message: "GetPonds_ID",
		Data:    result,
	}
}

func GetPonds_Ref(reference string, ref_id int, w http.ResponseWriter) Response_Ponds {
	connecting, err := connection(w)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer connecting.Close()
	data, err := connecting.Query("SELECT * FROM ponds WHERE "+reference+"=?", ref_id)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return Response_Ponds{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    []Ponds{},
		}
	}
	defer data.Close()

	var result []Ponds

	for data.Next() {
		var ponds = Ponds{}
		var err = data.Scan(&ponds.ID, &ponds.Farm_ID, &ponds.Name, &ponds.Thumbnails, &ponds.Description, &ponds.Created_at, &ponds.Updated_at, &ponds.Deleted_at)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return Response_Ponds{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
				Data:    []Ponds{},
			}
		}
		result = append(result, ponds)
	}

	return Response_Ponds{
		Status:  http.StatusOK,
		Message: "GetPonds_Ref",
		Data:    result,
	}
}
