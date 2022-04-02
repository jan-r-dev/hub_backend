package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

type article struct {
	Pk          int      `json:"pk,omitempty"`
	Title       string   `json:"title,omitempty"`
	Text        []string `json:"text,omitempty"`
	Image_url   []string `json:"image_url,omitempty"`
	Snippet_url []string `json:"snippet_url,omitempty"`
	Source_url  []string `json:"source_url,omitempty"`
}

type project struct {
	Pk          int       `json:"pk,omitempty"`
	Title       string    `json:"title,omitempty"`
	Summary     string    `json:"summary,omitempty"`
	Article_url string    `json:"article_url,omitempty"`
	Created_on  time.Time `json:"created_on,omitempty"`
	Stack       []string  `json:"stack,omitempty"`
}

func postgres(ctx context.Context, qs string) (pgx.Rows, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_DB"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		conn.Close(ctx)
		return nil, err
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, qs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while executing query: %v\n", err)
		rows.Close()
		return nil, err
	}

	return rows, err
}

func readRowsArticle(rows pgx.Rows) (article, error) {
	a := article{}
	for rows.Next() {
		errScan := rows.Scan(
			&a.Pk,
			&a.Title,
			&a.Text,
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

func readRowsProject(rows pgx.Rows) ([]project, error) {
	ps := []project{}
	for rows.Next() {
		p := project{}
		errScan := rows.Scan(
			&p.Pk,
			&p.Title,
			&p.Summary,
			&p.Article_url,
			&p.Created_on,
			&p.Stack,
		)
		if errScan != nil {
			fmt.Fprintf(os.Stderr, "Error while scanning rows in the Project table: %v\n", errScan)
			return ps, errScan
		}

		ps = append(ps, p)
	}

	return ps, nil
}
