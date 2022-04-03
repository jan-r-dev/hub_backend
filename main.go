package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	go getProjects(r)
	go getArticle(r)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run(":8080")
}

func getProjects(r *gin.Engine) {
	r.GET("/projects", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		ctx := context.Background()

		rows, err := postgres(
			ctx,
			`select project.pk, title, summary, article_url, created_date, stack from project
			inner join title on title_fk = title.pk
			inner join article on article_fk = article.pk
			order by pk asc;`,
		)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		}

		projects, err := readRowsProject(rows)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusOK, projects)
		}
	})
}

func getArticle(r *gin.Engine) {
	r.GET("/articles/:articleUrl", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		ctx := context.Background()
		articleUrl := "'" + c.Param("articleUrl") + "'"

		rows, err := postgres(
			ctx,
			(`select article.pk, title, text_content, image_url, snippet_url, source_url from article
			inner join title on article.title_fk = title.pk
			where article_url = ` + articleUrl),
		)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		}

		article, err := readRowsArticle(rows)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusOK, article)
		}
	})
}
