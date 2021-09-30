package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func connMongo(env map[string]string) {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + env["DB_USER"] + ":" + env["DB_PASS"] + "@cluster-1.eswlh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	fmt.Println(clientOptions)
}

/*

import "go.mongodb.org/mongo-driver/mongo"

clientOptions := options.Client().
    ApplyURI("mongodb+srv://jan_admin:<password>@cluster-1.eswlh.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}

*/
