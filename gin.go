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

		// for _, r := range projects {
		// 	fmt.Println(r.CreatedOn)
		// }

		c.JSON(http.StatusOK, projects)
	})

}
