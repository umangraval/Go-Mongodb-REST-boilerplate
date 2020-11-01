package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("POST: /person")
	response.Header().Add("context-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("golang").Collection("people")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(context.TODO(), person)
	json.NewEncoder(response).Encode(result)

}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
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

func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {
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

func DeletePersonEndpoint(response http.ResponseWriter, request *http.Request) {
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

func UpdatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
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

func UploadFileEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("POST: /upload")

	file, handler, err := request.FormFile("file")
	// fileName := request.FormValue("file_name")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile("uploaded/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.Copy(f, file)

	response.Write([]byte(`{ "message": "Uploaded Successfully" }`))
	json.NewEncoder(response)
}
