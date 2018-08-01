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

var book = model.Book{
	Title:    "Voyage au bout de la nuit",
	NumPages: 623,
	PubYear:  1930,
	AuthorID: "1234",
}

var author = model.Author{
	Name: "Celine",
}

type BookServerMock struct{}

func (bsm BookServerMock) CreateBook(book model.Book) (model.Book, error) {
	return book, nil
}

func (bsm BookServerMock) FindByTitle(title string) (model.Book, error) {
	return book, nil
}

func (bsm BookServerMock) List() ([]model.Book, error) {
	return []model.Book{book}, nil
}

type AuthorServerMock struct{}

func (asm AuthorServerMock) CreateAuthor(book model.Author) (model.Author, error) {
	return author, nil
}

func (asm AuthorServerMock) FindByID(ID string) (model.Author, error) {
	return author, nil
}

func (asm AuthorServerMock) FindByBookID(bookID string) (string, error) {
	return author.ID, nil
}

func TestBookQuery(t *testing.T) {
	logger := service.NewLogger(gcontext.LoadConfig())

	services := &service.Container{
		BookServer:   BookServerMock{},
		AuthorServer: AuthorServerMock{},
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "logger", logger)
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
						author {
							name
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"book": {
						"numPages": 623,
						"pubYear": 1930,
						"author": {
							"name": "Celine"
						}
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
						author {
							name
						}
					}
				}
			`,
			ExpectedResult: `
				{
					"books": [
						{
							"title": "Voyage au bout de la nuit",
							"author": {
								"name": "Celine"
							}
						}
					]
				}
			`,
		},
	})
}
