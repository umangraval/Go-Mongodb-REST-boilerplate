package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person Model
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func createPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("POST: /person")
	response.Header().Add("context-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("golang").Collection("people")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(context.TODO(), person)
	json.NewEncoder(response).Encode(result)

}

func getPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("GET: /people")
	response.Header().Add("content-type", "application/json")
	var people []*Person
	collection := client.Database("golang").Collection("people")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	for cursor.Next(context.TODO()) {
		var person Person
		err := cursor.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, &person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(people)
}

func getPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("context-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("GET: /person/" + params["id"])
	var person Person
	collection := client.Database("golang").Collection("people")
	err := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(person)

}

func deletePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("context-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("DELETE: /person/" + params["id"])
	collection := client.Database("golang").Collection("people")
	_, err := collection.DeleteOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{ "message": "Deleted" }`))
	json.NewEncoder(response)
}

func updatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("context-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("PUT: /person/" + params["id"])
	type fname struct {
		Firstname string `json:"firstname"`
	}
	var fir fname
	json.NewDecoder(request.Body).Decode(&fir)
	fmt.Println(fir.Firstname)
	collection := client.Database("golang").Collection("people")
	res, err := collection.UpdateOne(context.TODO(), bson.D{primitive.E{Key: "_id", Value: id}}, bson.D{{"$set", bson.D{{"firstname", fir.Firstname}}}})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(res)
}

func main() {
	fmt.Println("Server running on localhost:8080")
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
	client, _ = mongo.Connect(context.TODO(), clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/person", createPersonEndpoint).Methods("POST")
	router.HandleFunc("/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", deletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/person/{id}", updatePersonEndpoint).Methods("PUT")
	http.ListenAndServe(":8080", router)
}
