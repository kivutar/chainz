package resolver

import (
	"log"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/gqltesting"
	gcontext "github.com/kivutar/chainz/context"
	"github.com/kivutar/chainz/schema"
	"github.com/kivutar/chainz/service"
	"golang.org/x/net/context"
)

var (
	rootSchema = graphql.MustParseSchema(schema.GetRootSchema(), &Resolver{})
	ctx        context.Context
)

func init() {
	config := gcontext.LoadConfig()
	db, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	log := service.NewLogger(config)
	authorService := service.NewAuthorService(db, log)
	bookService := service.NewBookService(db, authorService, log)
	ctx = context.WithValue(context.Background(), "bookService", bookService)
}

func TestBasic(t *testing.T) {
	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Context: ctx,
			Schema:  rootSchema,
			Query: `
				{
					book(title:"The Holy Bible") {
						id
						numPages
					}
				}
			`,
			ExpectedResult: `
				{
					"book": {
					  "id": "1",
					  "numPages": 1000
					}
				}
			`,
		},
	})
}
