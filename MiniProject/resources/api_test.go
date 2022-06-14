package resources_test

import (
	"net/http"
	"strings"
	"testing"
)

func TestAPI_GET_Farm(t *testing.T) {
	reader := strings.NewReader(`{}`)
	request, err := http.NewRequest("GET", "http://localhost:9090/api/v1/farms", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_GET_Ponds(t *testing.T) {
	reader := strings.NewReader(`{}`)
	request, err := http.NewRequest("GET", "http://localhost:9090/api/v1/ponds", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_GET_APIAnalyst(t *testing.T) {
	reader := strings.NewReader(`{}`)
	request, err := http.NewRequest("GET", "http://localhost:9090/api/v1/api_analyst", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_GET_FarmByID(t *testing.T) {
	reader := strings.NewReader(`{}`)
	request, err := http.NewRequest("GET", "http://localhost:9090/api/v1/farms?id=1", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_GET_PondsByID(t *testing.T) {
	reader := strings.NewReader(`{}`)
	request, err := http.NewRequest("GET", "http://localhost:9090/api/v1/ponds?id=1", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_POST_Farm(t *testing.T) {
	reader := strings.NewReader(`{"Name":"Farm Name Unit Test 1",
                "Description":"Description Farm Name Unit Test 1",
                "Thumbnails":"Thumbnails Farm Name Unit Test 1"}`)
	request, err := http.NewRequest("POST", "http://localhost:9090/api/v1/farms", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_POST_Ponds(t *testing.T) {
	reader := strings.NewReader(`{"Farm_ID": 2,
                "Name":"Ponds Name Unit Test 1",
                "Description":"Description Pond Name Unit Test 1",
                "Thumbnails":"Thumbnails Pond Name Unit Test 1"}`)
	request, err := http.NewRequest("POST", "http://localhost:9090/api/v1/ponds", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_PUT_Farm(t *testing.T) {
	reader := strings.NewReader(`{"ID": 2,
                "Name":"Updating Farm Name Unit Test 2",
                "Description":"Updating Description Farm Name Unit Test 2",
                "Thumbnails":"Updating Thumbnails Farm Name Unit Test 2"}`)
	request, err := http.NewRequest("PUT", "http://localhost:9090/api/v1/farms", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_PUT_Ponds(t *testing.T) {
	reader := strings.NewReader(`{"ID": 2,
                "Farm_ID": 2,
                "Name":"Updating Ponds Name Unit Test 1",
                "Description":"Updating Description Pond Name Unit Test 1",
                "Thumbnails":"Updating Thumbnails Pond Name Unit Test 1"}`)
	request, err := http.NewRequest("PUT", "http://localhost:9090/api/v1/ponds", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_DELETE_Farm(t *testing.T) {
	reader := strings.NewReader(`{"ID": 2}`)
	request, err := http.NewRequest("DELETE", "http://localhost:9090/api/v1/farms", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}

func TestAPI_DELETE_Ponds(t *testing.T) {
	reader := strings.NewReader(`{"ID": 2}`)
	request, err := http.NewRequest("DELETE", "http://localhost:9090/api/v1/Ponds", reader)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		t.Error()
		return
	}
}
