package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := mongoConnClient(ctx)

	r := gin.New()

	getProjects(r, mongoClient)
	getArticle(r, mongoClient)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run()
}

func retrieveProjects(ts time.Time, mongoClient *mongo.Client) []Project {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoCollection := mongoAccessCollection("projects", mongoClient)

	//mongoFindProjects(ctx, mongoCollection)
	projects := mongoFindThreeProjects(ctx, mongoCollection, ts)

	return projects
}

func retrieveArticle(articleID string, mongoClient *mongo.Client) (Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoCollection := mongoAccessCollection("articles", mongoClient)

	article, err := mongoFindArticle(ctx, mongoCollection, articleID)

	return article, err
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
