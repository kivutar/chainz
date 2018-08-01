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
	authorService := ctx.Value("services").(*service.Container).AuthorServer
	logger := ctx.Value("logger").(*logging.Logger)

	author, err := authorService.FindByID(args.ID)
	if err != nil {
		logger.Errorf("Graphql error : %v", err)
		return nil, err
	}

	logger.Debugf("Retrieved author by author_id[%s] : %v", author.ID, author)
	return &AuthorResolver{&author}, nil
}
