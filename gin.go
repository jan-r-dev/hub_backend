package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getProjects(r *gin.Engine) {

	r.GET("/projects/:time", func(c *gin.Context) {
		ts := createTimestamp(c.Param("time"))

		projects := retrieveProjects(ts)

		// for _, r := range projects {
		// 	fmt.Println(r.CreatedOn)
		// }

		c.JSON(http.StatusOK, projects)
	})

}
