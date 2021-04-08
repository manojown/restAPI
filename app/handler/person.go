package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manojown/restApi/app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var total int = 0

func CreatePerson(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	var person model.Person
	err := json.NewDecoder(r.Body).Decode(person)
	if err != nil {
		ResponseWriter(w, http.StatusInternalServerError, "Error in send json", err.Error())
	}
	result, err := db.Collection("person").InsertOne(nil, person)
	if err != nil {
		ResponseWriter(w, http.StatusInternalServerError, "Error while creating person data", nil)
		return
	}
	person.ID = result.InsertedID.(primitive.ObjectID)
	ResponseWriter(w, http.StatusOK, "Person created successfully", person)

}

func GetPerson(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	var person model.Person

	err := db.Collection("person").FindOne(nil, model.Person{ID: id}).Decode(&person)
	if err != nil {
		ResponseWriter(w, http.StatusOK, err.Error(), nil)
		return
	}
	ResponseWriter(w, http.StatusOK, "Person get successfully.", person)

}

func GetTesting(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	// id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	var person model.Person

	person.Email = "manoj@gmail.com"
	person.FirstName = "manoj"
//	fmt.Println("Hit", total)
//	total++

	ResponseWriter(w, http.StatusOK, "Person get successfully.", person)

}
