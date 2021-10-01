package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// First step - Establish a client
func mongoConnClient() *mongo.Client {
	env := importEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+env["DB_USER"]+":"+env["DB_PASS"]+"@cluster-1.eswlh.mongodb.net/"+env["DB_DATABASE"]+"myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal("Error establishing Mongo client", err)
	}

	return client
}

// Second step - Access the chosen collection
func mongoAccessCollection(c string, client *mongo.Client) *mongo.Collection {
	collection := client.Database("hub").Collection(c)

	return collection
}

// Operation - Insert one into the collection specified in the argument
func mongoInsertOne(collection *mongo.Collection) {
	res, err := collection.InsertOne(context.Background(), bson.M{"hello1": "world2"})

	if err != nil {
		log.Fatal("Error inserting one into collection", err)
	}

	id := res.InsertedID

	fmt.Println(id)
}

// Operation - Find many in the collection according to the specified query
func mongoFindMany(collection *mongo.Collection, searchString string) {
	// To be continued ...
}
