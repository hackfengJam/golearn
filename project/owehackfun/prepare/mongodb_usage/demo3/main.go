package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.TODO(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
}
