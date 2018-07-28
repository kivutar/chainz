package resolver

import (
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/kivutar/chainz/model"
	"github.com/kivutar/chainz/service"
	"golang.org/x/net/context"
)

type BookResolver struct {
	u *model.Book
}

func (r *BookResolver) ID() graphql.ID {
	return graphql.ID(r.u.ID)
}

func (r *BookResolver) Title() string {
	return r.u.Title
}

func (r *BookResolver) NumPages() *int32 {
	return &r.u.NumPages
}

func (r *BookResolver) PubYear() *int32 {
	return &r.u.PubYear
}

func (r *BookResolver) Author(ctx context.Context) (*AuthorResolver, error) {
	author, err := ctx.Value("authorService").(*service.AuthorService).FindByID(r.u.AuthorID)
	return &AuthorResolver{
		author: author,
	}, err
}

func (r *BookResolver) CreatedAt() (*graphql.Time, error) {
	if r.u.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.u.CreatedAt)
	return &graphql.Time{Time: t}, err
}
