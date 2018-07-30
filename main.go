package main

import (
	"log"
	"net/http"

	gcontext "github.com/kivutar/chainz/context"
	h "github.com/kivutar/chainz/handler"
	"github.com/kivutar/chainz/resolver"
	"github.com/kivutar/chainz/schema"
	"github.com/kivutar/chainz/service"

	graphql "github.com/graph-gophers/graphql-go"
	"golang.org/x/net/context"
)

func main() {
	config := gcontext.LoadConfig()

	db, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	defer db.Close()

	ctx := context.Background()
	log := service.NewLogger(config)
	authorService := service.NewAuthorService(db, log)
	bookService := service.NewBookService(db, authorService, log)

	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "log", log)
	ctx = context.WithValue(ctx, "authorService", authorService)
	ctx = context.WithValue(ctx, "bookService", bookService)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	loggerHandler := &h.LoggerHandler{config.DebugMode}
	http.Handle("/query", h.AddContext(ctx, loggerHandler.Logging(&h.GraphQL{Schema: graphqlSchema})))

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
