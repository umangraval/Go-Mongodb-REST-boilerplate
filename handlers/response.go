package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/umangraval/Go-Mongodb-REST-boilerplate/models"
)

// AuthorizationResponse -> response authorize
func AuthorizationResponse(msg string, writer http.ResponseWriter) {
	type errdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &errdata{Statuscode: 401, Message: msg}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(writer).Encode(temp)
}

// SuccessArrRespond -> response formatter
func SuccessArrRespond(fields []*models.Person, writer http.ResponseWriter) {
	// var fields["status"] := "success"
	_, err := json.Marshal(fields)
	type data struct {
		People     []*models.Person `json:"data"`
		Statuscode int              `json:"status"`
		Message    string           `json:"msg"`
	}
	temp := &data{People: fields, Statuscode: 200, Message: "success"}
	if err != nil {
		ServerErrResponse(err.Error(), writer)
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// SuccessRespond -> response formatter
func SuccessRespond(fields models.Person, writer http.ResponseWriter) {
	_, err := json.Marshal(fields)
	type data struct {
		Person     models.Person `json:"data"`
		Statuscode int           `json:"status"`
		Message    string        `json:"msg"`
	}
	temp := &data{Person: fields, Statuscode: 200, Message: "success"}
	if err != nil {
		ServerErrResponse(err.Error(), writer)
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// SuccessResponse -> success formatter
func SuccessResponse(msg string, writer http.ResponseWriter) {
	type errdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &errdata{Statuscode: 200, Message: msg}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// ErrorResponse -> error formatter
func ErrorResponse(error string, writer http.ResponseWriter) {
	type errdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &errdata{Statuscode: 400, Message: error}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(temp)
}

// ServerErrResponse -> server error formatter
func ServerErrResponse(error string, writer http.ResponseWriter) {
	type servererrdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &servererrdata{Statuscode: 500, Message: error}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(temp)
}

// ValidationResponse -> user input validation
func ValidationResponse(fields map[string][]string, writer http.ResponseWriter) {
	//Create a new map and fill it
	response := make(map[string]interface{})
	response["errors"] = fields
	response["status"] = 422
	response["msg"] = "validation error"

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(writer).Encode(response)
}
