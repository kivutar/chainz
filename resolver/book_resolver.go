package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kivutar/chainz/model"
	"github.com/kivutar/chainz/service"
	"golang.org/x/net/context"
)

type BookResolver struct {
	book *model.Book
}

func (r *BookResolver) ID() graphql.ID {
	return graphql.ID(r.book.ID)
}

func (r *BookResolver) Title() string {
	return r.book.Title
}

func (r *BookResolver) NumPages() *int32 {
	return &r.book.NumPages
}

func (r *BookResolver) PubYear() *int32 {
	return &r.book.PubYear
}

func (r *BookResolver) Author(ctx context.Context) (*AuthorResolver, error) {
	author, err := ctx.Value("authorService").(*service.AuthorService).FindByID(r.book.AuthorID)
	return &AuthorResolver{
		author: author,
	}, err
}

func (r *BookResolver) CreatedAt() (*graphql.Time, error) {
	if r.book.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.book.CreatedAt)
	return &graphql.Time{Time: t}, err
}
