package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/umangraval/Go-Mongodb-REST-boilerplate/controllers"
)

func TestGetPeopleEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/people", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetPeopleEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"data":[{"_id":"5fa14f7485c5c7ef4c3c8f57","firstname":"umang","lastname":"raval"},{"_id":"5fa15187a0a10a088fb40401","firstname":"es","lastname":"raal"},{"_id":"5fa151fe0cabd3aeaf356280","firstname":"ess","lastname":"raaasl"},{"_id":"5fa1523477e6c0e0f8b69edf","firstname":"ess","lastname":"raaasl"},{"_id":"5fa152d01bc4dc5e90db219a","firstname":"ess","lastname":"raaasl"},{"_id":"5fa1530787b3a854d8abddf5","firstname":"ess","lastname":"raaasl"},{"_id":"5fa1b5dea1cc830278705074","lastname":"raaasl"}],"status":200,"msg":"success"}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}

func TestGetPersonEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/person/5fa14f7485c5c7ef4c3c8f57", nil)
	if err != nil {
		t.Fatal(err)
	}

	// req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetPersonEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected := `{"data": {"_id":"5fa14f7485c5c7ef4c3c8f57","firstname":"umang","lastname":"raval"},"status":200,"msg":"success"}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}

func TestCreatePersonEndpoint(t *testing.T) {

	var jsonStr = []byte(`{
		"firstname": "anto",
		"lastname": "manisha"
	}`)

	req, err := http.NewRequest("POST", "/person", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.CreatePersonEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// expected := `{
	// 	"status": 200,
	// 	"msg": "Inserted at 5fa4420153392120dee4e4d4"
	//   }`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}

func TestUpdatePersonEndpoint(t *testing.T) {
	var jsonStr = []byte(`{
		"firstname": "manoj"
	}`)
	req, err := http.NewRequest("PUT", "/person/5fa441ddf728403f388fa98f", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.UpdatePersonEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// expected := `{
	// 	"status": 200,
	// 	 "msg": "Updated"
	// }`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v",
	// 		rr.Body.String(), expected)
	// }
}
