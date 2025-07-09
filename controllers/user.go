package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/VineetBavniya/mongodb-golang.git/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	Collection *mongo.Collection
}


func NewUserController(client *mongo.Client) *UserController {
	coll := client.Database("mongo-golang").Collection("users")
	return &UserController{Collection: coll}
}

func (uc UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor , err := uc.Collection.Find(ctx, bson.M{})

	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(ctx)
	
	var users []models.User

	for cursor.Next(ctx){
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("⚠️ Error decoding user:", err)
			continue
		} 

		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, "Cursor error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}



func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	id := ps.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid, err := primitive.ObjectIDFromHex(id)
	
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	u := models.User{}
	
	err = uc.Collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u) 
	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}


func (uc UserController) CreateNewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	u.ID = primitive.NewObjectID()
	uc.Collection.InsertOne(context.TODO(), u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	
	res, err := uc.Collection.DeleteOne(context.TODO(), bson.M{"_id": oid})

	if err != nil || res.DeletedCount == 0 {
		http.Error(w, "Delete failed", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted User %s \n", oid)

}