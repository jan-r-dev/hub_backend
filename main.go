package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	myEnv := importEnv()
	fmt.Println(myEnv)

	connMongo(myEnv)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run()
}

/* NOTES
Mongo will be done using https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
Routing will be done using https://github.com/gin-gonic/gin#installation
Security will be done using


Mongo code:


1) Continue setting up the MongoDB connection in mongo.go

*/
