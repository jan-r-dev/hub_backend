package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient := mongoConnClient(ctx)

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

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		mongoCollection := mongoAccessCollection("projects", mongoClient)

		projects, err := mongoFindThreeProjects(ctx, mongoCollection, ts)

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
		mongoCollection := mongoAccessCollection("articles", mongoClient)

		article, err := mongoFindArticle(ctx, mongoCollection, articleId)

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
		mongoCollection := mongoAccessCollection("projects", mongoClient)

		projectCount, err := mongoCountProjects(ctx, mongoCollection)

		if err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusOK, projectCount)
		}
	})
}
