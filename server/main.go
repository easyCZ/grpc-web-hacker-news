package main

import (
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"google.golang.org/grpc"
	"net/http"
	"github.com/easyCZ/grpc-web-hacker-news/server/hackernews"
	hackernews_pb "github.com/easyCZ/grpc-web-hacker-news/server/proto"
	"google.golang.org/grpc/grpclog"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/go-chi/cors"
	"github.com/easyCZ/grpc-web-hacker-news/server/proxy"
	"github.com/easyCZ/grpc-web-hacker-news/server/middleware"
)

func main() {
	grpcServer := grpc.NewServer()
	hackernewsService := hackernews.NewHackerNewsService(nil)
	hackernews_pb.RegisterHackerNewsServiceServer(grpcServer, hackernewsService)

	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		middleware.NewGrpcWebMiddleware(wrappedGrpc).Handler,	// Must come before general CORS handling
		cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", },
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}).Handler,
	)

	router.Get("/article-proxy", proxy.Article)

	if err := http.ListenAndServe(":8900", router); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}
}
