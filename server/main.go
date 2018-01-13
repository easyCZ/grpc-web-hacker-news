package main

import (
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"net/http"
	"github.com/easyCZ/grpc-web-hacker-news/server/hackernews"
	hackernews_pb "github.com/easyCZ/grpc-web-hacker-news/server/proto"
)

func main() {
	router := chi.NewRouter()
	router.With()

	grpcServer := grpc.NewServer()
	hackernewsService := hackernews.New()
	hackernews_pb.RegisterHackerNewsServiceServer(grpcServer, hackernewsService)

	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	router.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})

	router.With(func(handler http.Handler) http.Handler {
		return wrappedGrpc
	})

	http.ListenAndServe(":8900", router)
}

