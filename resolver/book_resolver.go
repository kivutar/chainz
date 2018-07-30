package resolver

import (
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kivutar/chainz/model"
	"github.com/kivutar/chainz/service"
	logging "github.com/op/go-logging"
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
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	return &AuthorResolver{
		author: &author,
	}, nil
}

func (r *BookResolver) CreatedAt() *graphql.Time {
	return &graphql.Time{Time: r.book.CreatedAt}
}
