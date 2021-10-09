package main

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

var (
	db  *sql.DB
	err error
)
var client *mongo.Client
var lock sync.Mutex

type User struct {
	ID      int64  `json:"id"`
	Name   string `json:"name"`
	Email   string `json:"email"`
	Password int64  `json:"password"`
}
//Creating collection var as global variable for easy access across Functions
var collection = ConnecttoDB()



func ConnecttoDB() *mongo.Collection {

	// Set client options
	//change the URI according to your database
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.insta(), clientOptions)
	//Error Handling
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	//DB collection address which we are going to use
	//available to functions of all scope
	collection := client.Database("Appointy").Collection("users")

	return collection
}
// RequestHandlerFunc is the type defined to use the http Handler Function externally
type RequestHandlerFunc func(http.ResponseWriter, *http.Request)

// wrapper wraps the http request with sequence of middlewares provided
func wrapper(fn RequestHandlerFunc, mds ...func(RequestHandlerFunc) RequestHandlerFunc) RequestHandlerFunc {
	for _, md := range mds {
		fn = md(fn)
	}
	return fn
}
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
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(users)
}

func main(){
	
	fmt.Println("Starting the application...")
	fmt.Println(time.Now())
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
	defer client.Disconnect(ctx)
	
	handleRequests()
}

