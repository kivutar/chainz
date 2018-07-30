package resolver

import (
	"github.com/kivutar/chainz/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Author resolves an author graphql query
func (r *Resolver) Author(ctx context.Context, args struct {
	ID string
}) (*AuthorResolver, error) {
	author, err := ctx.Value("authorService").(*service.AuthorService).FindByID(args.ID)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	ctx.Value("log").(*logging.Logger).Debugf("Retrieved author by author_id[%s] : %v", author.ID, *author)
	return &AuthorResolver{author}, nil
}
