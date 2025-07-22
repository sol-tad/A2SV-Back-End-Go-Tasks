package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection

func ConnectDB(){
	// client,err:=mongo.NewClient(options.Client().ApplyURI("mongodb+srv://soltad65:12341234@taskmanagementcluster.tefktkx.mongodb.net/?retryWrites=true&w=majority&appName=TaskManagementCluster"))
	client,err:=mongo.NewClient(options.Client().ApplyURI("mongodb+srv://soltad65:12341234@taskmanagementcluster.tefktkx.mongodb.net/?retryWrites=true&w=majority&appName=TaskManagementCluster"))
	if err != nil {
		log.Fatal(err)
	}
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	err=client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	TaskCollection=client.Database("taskdb").Collection("tasks")
	UserCollection=client.Database("taskdb").Collection("users")
	log.Println("Connected to MongoDB")

}