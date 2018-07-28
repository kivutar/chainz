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

// // Books resolves a books graphql query
// func (r *Resolver) Books(ctx context.Context, args struct{}) (*booksConnectionResolver, error) {
// 	if isAuthorized := ctx.Value("is_authorized").(bool); !isAuthorized {
// 		return nil, errors.New(gcontext.CredentialsError)
// 	}
// 	bookID := ctx.Value("book_id").(*string)

// 	books, err := ctx.Value("bookService").(*service.BookService).List()
// 	count, err := ctx.Value("bookService").(*service.BookService).Count()
// 	ctx.Value("log").(*logging.Logger).Debugf("Retrieved books by book_id[%s] :", *bookID)
// 	config := ctx.Value("config").(*gcontext.Config)
// 	if config.DebugMode {
// 		for _, book := range books {
// 			ctx.Value("log").(*logging.Logger).Debugf("%v", *book)
// 		}
// 	}
// 	ctx.Value("log").(*logging.Logger).Debugf("Retrieved total books count by book_id[%s] : %v", *bookID, count)
// 	if err != nil {
// 		ctx.Value("log").(*logging.Logger).Errorf("Graphql error : %v", err)
// 		return nil, err
// 	}
// 	return &booksConnectionResolver{books: books, totalCount: count, from: &(books[0].ID), to: &(books[len(books)-1].ID)}, nil
// }
