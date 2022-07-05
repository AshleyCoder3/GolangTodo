package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AshleyCoder3/GolangTodo/graph"
	"github.com/AshleyCoder3/GolangTodo/graph/generated"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"time"
)

/* Tutorials I used in order
	1.  Build a Go Rest Api, React.js & TypeScript Todo Application.
		- https://www.youtube.com/watch?v=QevhhM_QfbM
 	2.	GraphQl in Go - GQLGen Tutorial
		- https://www.youtube.com/watch?v=O6jYy421tGw&t=783s
*/

const defaultPort = "8080"

func main() {
	//router := chi.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:4000"},
		//AllowedHeaders:   []string{"Origin, Content-Type, Accept"},
		AllowCredentials: true,
		//Debug:            true,
	})

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//srv.AddTransport(transport.POST{})
	//srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	//srv.Use(extension.Introspection{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))
	http.Handle("/api/todos", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	//err := http.ListenAndServe(":8080", router)
	//if err != nil {
	//	panic(err)
	//}

}
