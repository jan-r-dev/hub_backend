package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	getProjects(r)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run()
}

func retrieveProjects(ts time.Time) []Project {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	mongoClient := mongoConnClient(ctx)

	mongoCollection := mongoAccessCollection("projects", mongoClient)

	//mongoFindProjects(ctx, mongoCollection)
	projects := mongoFindThreeProjects(ctx, mongoCollection, ts)

	return projects
}

func createTimestamp(stringTime string) time.Time {
	i, err := strconv.ParseInt(stringTime, 10, 64)

	if err != nil {
		log.Fatal("Error conversion: ", err)
	}
	tm := time.Unix(i, 0)
	//fmt.Println(tm)

	return tm
}

//1633698368

/* NOTES
Mongo will be done using https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo
Routing will be done using https://github.com/gin-gonic/gin#installation
Security will be done using lord only knows what
*/
