package resolver

import (
	"github.com/kivutar/chainz/model"
	"github.com/kivutar/chainz/service"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)

// CreateBook creates a book record in the database
func (r *Resolver) CreateBook(ctx context.Context, args *struct {
	Title    string
	PubYear  *int32
	NumPages *int32
	AuthorID *string
}) (*BookResolver, error) {
	book := model.Book{
		Title: args.Title,
	}
	if args.PubYear != nil {
		book.PubYear = *args.PubYear
	}
	if args.NumPages != nil {
		book.NumPages = *args.NumPages
	}
	if args.AuthorID != nil {
		book.AuthorID = *args.AuthorID
	}

	book, err := ctx.Value("bookService").(*service.BookService).CreateBook(book)
	if err != nil {
		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
		return nil, err
	}
	ctx.Value("log").(*logging.Logger).Debugf("Created book : %v", book)
	return &BookResolver{&book}, nil
}
