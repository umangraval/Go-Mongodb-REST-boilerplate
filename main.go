package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/controllers"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	fmt.Println("Server running on localhost:8080")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

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
	http.ListenAndServe(":8080", logRequest(handler))
}
