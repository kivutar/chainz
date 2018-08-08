package resolver

import (
	"github.com/kivutar/chainz/model"
	"github.com/kivutar/chainz/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateAuthor creates an author record in the database
func (r *Resolver) CreateAuthor(ctx context.Context, args *struct {
	Name string
}) (*AuthorResolver, error) {
	authorService := ctx.Value("services").(*service.Container).AuthorServer
	logger := ctx.Value("logger").(*logging.Logger)

	author := model.Author{
		Name: args.Name,
	}

	author, err := authorService.CreateAuthor(author)
	if err != nil {
		logger.Errorf("Graphql error : %v", err)
		return nil, err
	}
	logger.Debugf("Created author : %v", author)
	return &AuthorResolver{&author}, nil
}
