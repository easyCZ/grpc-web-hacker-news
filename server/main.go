package main

import (
	"log"
	"net/http"

	"github.com/easyCZ/grpc-web-hacker-news/server/hackernews"
	"github.com/easyCZ/grpc-web-hacker-news/server/middleware"
	hackernews_pb "github.com/easyCZ/grpc-web-hacker-news/server/proto"
	"github.com/easyCZ/grpc-web-hacker-news/server/proxy"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	hackernewsService := hackernews.NewHackerNewsService(nil)
	hackernews_pb.RegisterHackerNewsServiceServer(grpcServer, hackernewsService)

	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithOriginFunc(func(origin string) bool {
		// Allow all origins, DO NOT do this in production
		return true
	}))

	router := chi.NewRouter()
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		middleware.NewGrpcWebMiddleware(wrappedGrpc).Handler,
	)

	router.Get("/article-proxy", proxy.Article)

	log.Println("Serving API on http://127.0.0.1:8900")
	if err := http.ListenAndServe(":8900", router); err != nil {
		log.Fatalf("failed starting http2 server: %v", err)
	}
}
