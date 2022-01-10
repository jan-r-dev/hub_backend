package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Project struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Summary   string             `json:"summary,omitempty" bson:"summary,omitempty"`
	Stack     []string           `json:"stack,omitempty" bson:"stack,omitempty"`
	Link      string             `json:"link,omitempty" bson:"link,omitempty"`
	CreatedOn time.Time          `json:"created_on,omitempty" bson:"created_on,omitempty"`
	ArticleId primitive.ObjectID `json:"article_id,omitempty" bson:"article_id"`
}

type Article struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	Images   []string           `json:"images,omitempty" bson:"images,omitempty"`
	Snippets []string           `json:"snippets,omitempty" bson:"snippets,omitempty"`
	Text     []string           `json:"text,omitempty" bson:"text,omitempty"`
	Sources  []string           `json:"sources,omitempty" bson:"sources,omitempty"`
	Link     string             `json:"link,omitempty" bson:"link,omitempty"`
}

// First step - Establish a client
func mongoConnClient(ctx context.Context) *mongo.Client {
	env := importEnv()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+env["DB_USER"]+":"+env["DB_PASS"]+"@cluster-1.eswlh.mongodb.net/"+env["DB_DATABASE"]+"myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal("Error establishing Mongo client", err)
	}

	return client
}

// Second step - Access the chosen collection
func mongoAccessCollection(c string, client *mongo.Client) *mongo.Collection {
	collection := client.Database("hub").Collection(c)

	return collection
}

// Operation - Find and return all projects sorted by date (descending)
//lint:ignore U1000 placeholder
func mongoFindProjects(ctx context.Context, collection *mongo.Collection) {

	options := options.Find().SetSort(bson.D{primitive.E{Key: "created_on", Value: -1}})

	cursor, err := collection.Find(ctx, bson.D{
		primitive.E{},
	}, options)
	if err != nil {
		log.Fatal("Error searching the collection: ", err)
	}

	var projects []Project

	for cursor.Next(ctx) {
		var project Project

		cursor.Decode(&project)
		projects = append(projects, project)
	}

	for _, r := range projects {
		fmt.Println(r.CreatedOn)
	}
}

func mongoFindThreeProjects(ctx context.Context, collection *mongo.Collection, startPoint time.Time) []Project {
	options := options.Find().SetSort(bson.D{primitive.E{Key: "created_on", Value: -1}})

	cursor, err := collection.Find(ctx, bson.D{
		primitive.E{Key: "created_on", Value: bson.D{
			primitive.E{Key: "$lt", Value: startPoint},
		},
		},
	}, options)

	if err != nil {
		log.Fatal("Error searching the collection: ", err)
	}

	var projects []Project

	counter := 0
	for cursor.Next(ctx) && counter != 3 {

		var project Project

		cursor.Decode(&project)
		projects = append(projects, project)

		counter += 1
	}

	return projects
}

func mongoFindArticle(ctx context.Context, collection *mongo.Collection, articleId string) (Article, error) {
	singleResult := collection.FindOne(ctx, bson.D{
		primitive.E{Key: "link", Value: articleId},
	})

	var article Article

	err := singleResult.Decode(&article)

	return article, err
}

/*
// MISC

// Operation - Insert one into the collection specified in the argument
func mongoInsertOne(collection *mongo.Collection, entryPair struct{}) {
	res, err := collection.InsertOne(context.Background(), bson.M{"hello1": "world2"})

	if err != nil {
		log.Fatal("Error inserting one into collection", err)
	}

	id := res.InsertedID

	fmt.Println(id)
}


*/
