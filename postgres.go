package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

type HubAPI interface {
	endpoint()
}

type article struct {
	Pk           int      `json:"pk,omitempty"`
	Title        string   `json:"title,omitempty"`
	Text_content []string `json:"text,omitempty"`
	Image_url    []string `json:"images,omitempty"`
	Snippet_url  []string `json:"snippets,omitempty"`
	Source_url   []string `json:"sources,omitempty"`
}

type project struct {
	Pk         int       `json:"pk,omitempty"`
	Title      string    `json:"title,omitempty"`
	Summary    string    `json:"summary,omitempty"`
	ArticleUrl string    `json:"articleurl,omitempty"`
	Created_on time.Time `json:"created_on,omitempty"`
	Stack      []string  `json:"stack,omitempty"`
}

func (a article) endpoint() {

}

func (p project) endpoint() {
}

func postgres(ctx context.Context, qs string, readRows func(pgx.Rows) (HubAPI, error)) (HubAPI, error) {
	conn, errConn := pgx.Connect(ctx, os.Getenv("POSTGRES_DB"))
	if errConn != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", errConn)
		conn.Close(ctx)
		return article{}, errConn
	}
	defer conn.Close(ctx)

	rows, errQuery := conn.Query(ctx, qs)
	if errQuery != nil {
		fmt.Fprintf(os.Stderr, "Error while executing query: %v\n", errQuery)
		rows.Close()
		return article{}, errQuery
	}
	defer rows.Close()

	return readRows(rows)
}

func readRowsArticle(rows pgx.Rows) (HubAPI, error) {
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
			fmt.Fprintf(os.Stderr, "Error while scanning rows in the Article table: %v\n", errScan)
			return a, errScan
		}
	}

	return a, nil
}

func readRowsProject(rows pgx.Rows) (HubAPI, error) {
	p := project{}
	for rows.Next() {
		errScan := rows.Scan(
			&p.Pk,
			&p.Title,
			&p.Summary,
			&p.ArticleUrl,
			&p.Created_on,
			&p.Stack,
		)
		if errScan != nil {
			fmt.Fprintf(os.Stderr, "Error while scanning rows in the Project table: %v\n", errScan)
			return p, errScan
		}
	}

	return p, nil
}
