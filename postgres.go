package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type article struct {
	Pk           int      `json:"pk,omitempty"`
	Title        string   `json:"title,omitempty"`
	Text_content []string `json:"text,omitempty"`
	Image_url    []string `json:"images,omitempty"`
	Snippet_url  []string `json:"snippets,omitempty"`
	Source_url   []string `json:"sources,omitempty"`
}

type request struct {
	endpoint string
	id       int
}

func postgres(ctx context.Context, id string) (article, error) {
	conn, errConn := pgx.Connect(ctx, os.Getenv("POSTGRES_DB"))
	if errConn != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", errConn)
		conn.Close(ctx)
		return article{}, errConn
	}
	defer conn.Close(ctx)

	qs := "select article.pk, title, text_content, image_url, snippet_url, source_url from article inner join title on article.title_fk = title.pk where article.pk = " + id

	rows, errQuery := conn.Query(ctx, qs)
	if errQuery != nil {
		fmt.Fprintf(os.Stderr, "Error while executing query: %v\n", errQuery)
		rows.Close()
		return article{}, errQuery
	}
	defer rows.Close()

	return readRowsArticle(rows)
}

func readRowsArticle(rows pgx.Rows) (article, error) {
	a := article{}
	for rows.Next() {

		errScan := rows.Scan(
			&a.Pk,
			&a.Title,
			&a.Text_content,
			&a.Image_url,
			&a.Snippet_url,
			&a.Source_url,
		)
		if errScan != nil {
			fmt.Fprintf(os.Stderr, "Error while scanning rows: %v\n", errScan)
			return a, errScan
		}
	}

	return a, nil
}
