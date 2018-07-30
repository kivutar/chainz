package resolver

import (
	"log"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
	gcontext "github.com/kivutar/chainz/context"
	"github.com/kivutar/chainz/model"
	"github.com/kivutar/chainz/schema"
	"github.com/kivutar/chainz/service"
	"golang.org/x/net/context"
)

var (
	graphqlSchema *graphql.Schema
	ctx           context.Context
)

func TestBookQuery(t *testing.T) {
	config := gcontext.LoadConfig()

	db, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	defer db.Close()

	ctx := context.Background()
	log := service.NewLogger(config)
	authorService := service.NewAuthorService(db, log)
	bookService := service.NewBookService(db, authorService, log)

	ctx = context.WithValue(ctx, "log", log)
	ctx = context.WithValue(ctx, "bookService", bookService)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &Resolver{})

	bookService.CreateBook(model.Book{
		Title:    "Voyage au bout de la nuit",
		NumPages: 623,
		PubYear:  1930,
		AuthorID: "1234",
	})

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  graphqlSchema,
			Query: `
				{
					book(title: "Voyage au bout de la nuit") {
						numPages
						pubYear
					}
				}
			`,
			ExpectedResult: `
				{
					"book": {
						"numPages": 623,
						"pubYear": 1930
					}
				}
			`,
		},
		{
			Context: ctx,
			Schema:  graphqlSchema,
			Query: `
				{
					books {
						title
					}
				}
			`,
			ExpectedResult: `
				{
					"books": [
						{
							"title": "Voyage au bout de la nuit"
						}
					]
				}
			`,
		},
	})
}
