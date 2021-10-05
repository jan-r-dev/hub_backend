package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func mongoInsertOne(collection *mongo.Collection, entryPair struct{}) {
	res, err := collection.InsertOne(context.Background(), bson.M{"hello1": "world2"})

	if err != nil {
		log.Fatal("Error inserting one into collection", err)
	}

	id := res.InsertedID

	fmt.Println(id)
}

// Operation - Find and return all projects sorted by date (descending)

func mongoFindProjects(collection *mongo.Collection, timestamp string) *mongo.Cursor {
	// filter by created_on
	// -1 marks descending, 1 ascending
	options := options.Find().SetSort(bson.D{primitive.E{Key: "created_on", Value: -1}})

	// Next step - modify the filter
	cursor, err := collection.Find(context.Background(), bson.D{
		{"created_on": {""}}}, options)
	if err != nil {
		log.Fatal("Error searching the collection: ", err)
	}

	return cursor
}

// Cursor - Print all results
func mongoPrintCursor(c *mongo.Cursor) {
	var results []bson.M
	if err := c.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

/*
func e() {
	var coll *mongo.Collection

	// Find all documents in which the "name" field is "Bob".
	// Specify the Sort option to sort the returned documents by age in
	// ascending order.
	opts := options.Find().SetSort(bson.D{{"aged", 1}})
	cursor, err := coll.Find(context.TODO(), bson.D{{"name", "Bob"}}, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all returned documents and print them out.
	// See the mongo.Cursor documentation for more examples of using cursors.
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
*/
