package resolver

import (
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

type BookServerMocker struct{}

func (bfm BookServerMocker) FindByTitle(title string) (model.Book, error) {
	return model.Book{
		Title:    "Voyage au bout de la nuit",
		NumPages: 623,
		PubYear:  1930,
		AuthorID: "1234",
	}, nil
}

func (bfm BookServerMocker) List() ([]model.Book, error) {
	return []model.Book{
		model.Book{
			Title:    "Voyage au bout de la nuit",
			NumPages: 623,
			PubYear:  1930,
			AuthorID: "1234",
		},
	}, nil
}

func TestBookQuery(t *testing.T) {
	config := gcontext.LoadConfig()

	log := service.NewLogger(config)
	bookService := BookServerMocker{}

	services := &service.Container{
		BookServer: bookService,
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "log", log)
	ctx = context.WithValue(ctx, "services", services)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &Resolver{})

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
