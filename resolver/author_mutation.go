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
	author := model.Author{
		Name: args.Name,
	}

	author, err := ctx.Value("services").(*service.Container).AuthorServer.CreateAuthor(author)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created author : %v", author)
	return &AuthorResolver{&author}, nil
}
