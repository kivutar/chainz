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
	book, err := ctx.Value("bookService").(*service.BookService).FindByTitle(args.Title)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Retrieved book by book_id[%s] : %v", book.ID, *book)
	return &BookResolver{book}, nil
}

// Books resolves a books graphql query
func (r *Resolver) Books(ctx context.Context) (*[]*BookResolver, error) {
	books, err := ctx.Value("bookService").(*service.BookService).List()
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	brs := make([]*BookResolver, 0)
	for _, book := range books {
		brs = append(brs, &BookResolver{
			book: book,
		})
	}
	return &brs, nil
}
