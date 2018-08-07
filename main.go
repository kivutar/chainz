package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	gcontext.InitSessionStore()

	db, err := gcontext.OpenDB(config)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s \n", err)
	}
	defer db.Close()

	ctx := context.Background()
	logger := service.NewLogger(config)
	authorService := service.NewAuthorService(db)
	bookService := service.NewBookService(db)

	services := &service.Container{
		BookServer:   bookService,
		AuthorServer: authorService,
	}

	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "logger", logger)
	ctx = context.WithValue(ctx, "services", services)

	graphqlSchema := graphql.MustParseSchema(schema.GetRootSchema(), &resolver.Resolver{})

	r := mux.NewRouter()

	loggerHandler := &h.LoggerHandler{DebugMode: config.DebugMode}
	r.HandleFunc("/query", h.AddContext(ctx, loggerHandler.Logging(&h.GraphQL{Schema: graphqlSchema})))

	r.HandleFunc("/callback", h.CallbackHandler)

	r.HandleFunc("/login", h.LoginHandler)

	r.HandleFunc("/user", h.UserHandler)

	r.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "graphiql.html")
	}))

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
