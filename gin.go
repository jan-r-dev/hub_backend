package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func getProjects(r *gin.Engine, mongoClient *mongo.Client) {

	r.GET("/projects/:time", func(c *gin.Context) {
		ts := createTimestamp(c.Param("time"))

		projects := retrieveProjects(ts, mongoClient)

		c.JSON(http.StatusOK, projects)
	})

}

func getArticle(r *gin.Engine, mongoClient *mongo.Client) {
	r.GET("/articles/:articleId", func(c *gin.Context) {
		articleId := c.Param("articleId")

		article := retrieveArticle(articleId, mongoClient)
		c.JSON(http.StatusOK, article)
	})
}

/*
	articleId, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		panic(err)
	}
*/
