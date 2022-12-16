package model

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var MongoClient = InitMongoClient() 

func InitMongoClient() *mongo.Database {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "admin",
	}).ApplyURI("mongodb://172.17.0.2:27017"))
	
	if err != nil {
		log.Println("连接 Mongo 错误"+ err.Error())
		return nil
	}
	
	db := client.Database("im")	
	
	return db
}
