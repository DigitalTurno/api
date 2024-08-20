package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/diegofly91/apiturnos/src/generated"
	"github.com/diegofly91/apiturnos/src/module"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	res := module.SetupModule()
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: res}))
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		var myErr *gqlerror.Error
		if errors.As(e, &myErr) {
			err.Message = myErr.Message
		}

		return err
	})
	http.Handle("/", playground.Handler("Api DigitalTurno GraphQL", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
