package utils


import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)



//create an user
func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("milhub_user").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

// GetUser returns a user
func GetUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("context-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user User
	collection := client.Database("milhub_user").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, User{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
	}
	json.NewEncoder(response).Encode(user)
}


func GetUsersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var users []User
	collection := client.Database("milhub_user").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}