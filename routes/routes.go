package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/controllers"
	middlewares "github.com/umangraval/Go-Mongodb-REST-boilerplate/handlers"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/person", controllers.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/auth", controllers.Auths).Methods("GET")
	router.HandleFunc("/people", middlewares.IsAuthorized(controllers.GetPeopleEndpoint)).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", controllers.DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/person/{id}", controllers.UpdatePersonEndpoint).Methods("PUT")
	router.HandleFunc("/upload", controllers.UploadFileEndpoint).Methods("POST")
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./uploaded/"))))
	return router
}
