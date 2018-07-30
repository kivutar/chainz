package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/kivutar/chainz/model"
)

type AuthorResolver struct {
	author *model.Author
}

func (r *AuthorResolver) ID() graphql.ID {
	return graphql.ID(r.author.ID)
}

func (r *AuthorResolver) Name() string {
	return r.author.Name
}

func (r *AuthorResolver) CreatedAt() *graphql.Time {
	return &graphql.Time{Time: r.author.CreatedAt}
}
