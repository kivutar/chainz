package resolver

import (
	"github.com/kivutar/chainz/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// Book resolves a book graphql query
func (r *Resolver) Book(ctx context.Context, args struct {
	Title string
}) (*BookResolver, error) {
	bookService := ctx.Value("services").(*service.Container).BookServer
	logger := ctx.Value("logger").(*logging.Logger)

	book, err := bookService.FindByTitle(args.Title)
	if err != nil {
		logger.Errorf("Graphql error : %v", err)
		return nil, err
	}

	logger.Debugf("Retrieved book by title[%s] : %v", args.Title, book)
	return &BookResolver{&book}, nil
}

// Books resolves a books graphql query
func (r *Resolver) Books(ctx context.Context) (*[]*BookResolver, error) {
	bookService := ctx.Value("services").(*service.Container).BookServer

	books, err := bookService.List()
	if err != nil {
		//ctx.Value("logger").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}

	var brs []*BookResolver
	for _, book := range books {
		brs = append(brs, &BookResolver{
			book: &book,
		})
	}
	return &brs, nil
}
