package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vspirandeli/fullcycle-3-0-mod-4-graphql/graph"
	"github.com/vspirandeli/fullcycle-3-0-mod-4-graphql/internal/database"
)

const defaultPort = "8080"

func main() {
	db, error := sql.Open("sqlite3", "./data.db")
	if error != nil {
		log.Fatalf("failed to open database: %v", error)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{CategoryDB: categoryDb, CourseDB: courseDb}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
