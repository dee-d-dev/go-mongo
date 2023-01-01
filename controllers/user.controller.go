package controllers

import (
	"encoding/json"
// 	"fmt"
	"net/http"
	"context"
	"time"

	"github.com/dee-d-dev/go-mongo/models"
// 	"github.com/julienschmidt/httprouter"
// 	"gopkg.in/mgo.v2/bson"
// 	"gopkg.in/mgo.v2"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// type UserController struct {
// 	session *mgo.Session
// }

var client *mongo.Client

func CreatePerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var person models.User
	json.NewDecoder(r.Body).Decode(&person)
	collection := client.Database("g0-mongo").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(w).Encode(result)

}