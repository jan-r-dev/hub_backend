package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	mongoClient := mongoConnClient()

	// Specify collection name
	mongoCollection := mongoAccessCollection("projects", mongoClient)

	// Retrieve cursor
	mongoCursor := mongoFindProjects(mongoCollection)

	mongoPrintCursor(mongoCursor)

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
