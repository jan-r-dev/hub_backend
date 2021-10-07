package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	mongoClient := mongoConnClient(ctx)

	mongoCollection := mongoAccessCollection("projects", mongoClient)

	mongoFindProjects(ctx, mongoCollection)

	// r := gin.New()

	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())

	// r.Run()
}

/* NOTES
Mongo will be done using https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
Routing will be done using https://github.com/gin-gonic/gin#installation
Security will be done using lord only knows what
*/
