package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/diegofly91/apiturnos/src"
	"github.com/diegofly91/apiturnos/src/modules/auth/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.Use(middleware.AuthMiddleware)
	srv := src.AppHandles()
	router.Handle("/", playground.Handler("Api DigitalTurno GraphQL", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
