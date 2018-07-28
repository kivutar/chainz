package resolver

import (
	"time"

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

func (r *AuthorResolver) CreatedAt() (*graphql.Time, error) {
	if r.author.CreatedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.author.CreatedAt)
	return &graphql.Time{Time: t}, err
}
