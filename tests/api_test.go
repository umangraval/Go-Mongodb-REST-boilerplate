package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/umangraval/Go-Mongodb-REST-boilerplate/routes"
)

var rou = routes.Routes()

func TestGetPeopleEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/people", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	rou.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if !strings.Contains(rr.Body.String(), "success") {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}

func TestGetPersonEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/person/5fa1530787b3a854d8abddf5", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	rou.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if !strings.Contains(rr.Body.String(), "success") {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}

func TestCreatePersonEndpoint(t *testing.T) {
	var jsonStr = []byte(`{
		"firstname": "anto",
		"lastname": "karis"
	}`)

	req, err := http.NewRequest("POST", "/person", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	rou.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if !strings.Contains(rr.Body.String(), "Inserted") {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}

func TestUpdatePersonEndpoint(t *testing.T) {
	var jsonStr = []byte(`{
		"firstname": "manoj"
	}`)
	req, err := http.NewRequest("PUT", "/person/5fa1530787b3a854d8abddf5", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	rou.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if !strings.Contains(rr.Body.String(), "Updated") {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}

func TestDeletePersonEndpoint(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/person/5fa1530787b3a854d8abddf5", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	rou.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if !strings.Contains(rr.Body.String(), "Deleted") {
		t.Errorf("handler returned unexpected body: got %v",
			rr.Body.String())
	}
}
