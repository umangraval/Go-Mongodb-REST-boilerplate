package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/controllers"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	fmt.Println("Server running on localhost:8080")
	router := mux.NewRouter()

	router.HandleFunc("/person", controllers.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people", controllers.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/person/{id}", controllers.UpdatePersonEndpoint).Methods("PUT")
	router.HandleFunc("/upload", controllers.UploadFileEndpoint).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(":8080", handler)
}
