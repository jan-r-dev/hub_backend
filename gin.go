package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getProjects(r *gin.Engine, mongoClient *mongo.Client) {

	r.GET("/projects/:time", func(c *gin.Context) {
		ts := createTimestamp(c.Param("time"))

		projects := retrieveProjects(ts, mongoClient)

		// for _, r := range projects {
		// 	fmt.Println(r.CreatedOn)
		// }

		c.JSON(http.StatusOK, projects)
	})

}

func getArticle(r *gin.Engine, mongoClient *mongo.Client) {
	r.GET("/articles/:id", func(c *gin.Context) {
		s := c.Param("id")

		articleID, err := primitive.ObjectIDFromHex(s)
		if err != nil {
			log.Fatal(err)
		}

		article := retrieveArticle(articleID, mongoClient)
		c.JSON(http.StatusOK, article)
	})
}

/*

func ObjectIDFromHex(s string) (ObjectID, error)

Use it as follows:

objID, err := primitive.ObjectIDFromHex(hexString)
if err != nil {
  panic(err)
}

*/
