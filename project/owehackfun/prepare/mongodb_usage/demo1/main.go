package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
	)
	// 1. 建立连接
	/*
		if client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017")); err != nil {
			fmt.Println(err)
			return
		}
		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	*/
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017")); err != nil {
		fmt.Println(err)
		return
	}
	// 2. 选择数据库
	database = client.Database("test")

	// 3. 选择表 my_collection
	collection = database.Collection("my_collection")

	collection = collection

}
