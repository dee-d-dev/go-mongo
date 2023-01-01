package main

import (
	"context"
	"encoding/json"

	// "log"
	"net/http"

	// "os"
	"fmt"
	"time"

	// "github.com/dee-d-dev/go-mongo/controllers"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/gorilla/mux"
)

type User struct {
	// ID        primitive.ObjectID `json:"_id,omitempty" bson: "id,omitempty"`
	firstname string `json: "firstname,omitempty" bson:"firstname,omitempty"`
	lastname  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func CreateUser(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("content-type", "application/json")
	var user User

	json.NewDecoder(req.Body).Decode(&user)
	collection := client.Database("go-crud").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(resp).Encode(result)
}

func GetUsers(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("content-type", "application/json")
	var users []User
	collection := client.Database("go-crud").Collection("Users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"message": "` + err.Error() + ` "}`))
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user User
		cursor.Decode(&user)
		json.NewEncoder(resp).Encode(users)
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"message": "` + err.Error() + ` "}`))
		return
	}

	// json.NewEncoder(resp).Encode(users)
}

func main() {
	fmt.Println("starting application")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://admin:dbpassword@cluster0.30nkvtf.mongodb.net/?retryWrites=true&w=majority"))
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:dbpassword@cluster0.30nkvtf.mongodb.net/?retryWrites=true&w=majority"))

	router := mux.NewRouter()
	router.HandleFunc("/user", CreateUser).Methods("POST")
	router.HandleFunc("/users", GetUsers).Methods("GET")

	http.ListenAndServe(":5060", router)
}

//mongodb+srv://admin:admin@cluster0.30nkvtf.mongodb.net/?retryWrites=true&w=majority
// mongodb://localhost:27017
