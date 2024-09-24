package main

import (
	"log"
	"net/http"
	"os"

	"apiturnos/src"
	"apiturnos/src/modules/auth/middleware"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Ajusta para tu frontend
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})
	router := mux.NewRouter()
	authMiddleware := middleware.NewAuthMiddleware()
	router.Use(authMiddleware.Auth)
	srv := src.PlaygroundHandler()
	router.Handle("/", playground.ApolloSandboxHandler("Api DigitalTurno GraphQL", "/query"))
	// Iniciar el servidor con el middleware CORS aplicado correctamente
	router.Handle("/query", c.Handler(srv))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
