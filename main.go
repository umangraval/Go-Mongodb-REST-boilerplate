package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	fmt.Println("Server running on localhost:8080")
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, _ = mongo.Connect(context.TODO(), clientOptions)

	router := mux.NewRouter()
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// originsOk := handlers.AllowedOrigins([]string{"https://prog.in"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/person/{id}", UpdatePersonEndpoint).Methods("PUT")
	router.HandleFunc("/upload", UploadFileEndpoint).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(":8080", handler)
}
