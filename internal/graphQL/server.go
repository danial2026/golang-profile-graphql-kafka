package main

import (
	"log"
	"net/http"
	"os"

	"github.com/danial2026/golang-profile-graphql-kafka/internal/graphql/graph"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/graphql/client"
	"github.com/danial2026/golang-profile-graphql-kafka/internal/graphql/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	client.ConnectToService()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
