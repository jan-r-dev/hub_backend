package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := connectClientDB(ctx)

	r := gin.New()

	go getProjects(r, mongoClient)
	go getArticle(r, mongoClient)
	go getProjectCount(r, mongoClient)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run(":8080")
}

func getProjects(r *gin.Engine, mongoClient *mongo.Client) {
	r.GET("/projects/:time", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		ts := createTimestamp(c.Param("time"))
		count, _ := strconv.ParseInt(c.Query("count"), 10, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		mongoCollection := accessCollectionDB("projects", mongoClient)

		projects, err := getProjectsFromDB(ctx, mongoCollection, ts, count)

		if err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusOK, projects)
		}
	})

}

func getArticle(r *gin.Engine, mongoClient *mongo.Client) {
	r.GET("/articles/:articleId", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		articleId := c.Param("articleId")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		mongoCollection := accessCollectionDB("articles", mongoClient)

		article, err := getArticleFromDB(ctx, mongoCollection, articleId)

		if err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusOK, article)
		}
	})
}

func getProjectCount(r *gin.Engine, mongoClient *mongo.Client) {
	r.GET("/projects", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		mongoCollection := accessCollectionDB("projects", mongoClient)

		projectCount, err := getProjectCountFromDB(ctx, mongoCollection)

		if err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusOK, projectCount)
		}
	})
}
