package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/VineetBavniya/mongodb-golang.git/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getSeason() *mongo.Client{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := "mongodb://root:root123@localhost:27017/?authSource=admin" // put here you url 
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("MongoDB connect error:", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping error:", err)
	}
	
	log.Println("✅ Connected to MongoDB")

	return client
}




func main(){

	router := httprouter.New()

	userController := controllers.NewUserController(getSeason())

	router.GET("/user", userController.GetAllUsers)
	router.GET("/user/:id",	userController.GetUser)
	router.POST("/user", userController.CreateNewUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	http.ListenAndServe(":9001", router)
}

