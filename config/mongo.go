package config

import (
	"context"
	"log"
	"time"

	"employees/controller"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDB() {
	// Database Config
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected to Database!")
	}
	db := client.Database("employees")
	controller.EmployeeCollection(db)
	return
}
