package main

import (
	"github.com/go-chi/chi"
	"google.golang.org/grpc"
	"net/http"
	"github.com/easyCZ/grpc-web-hacker-news/server/hackernews"
	hackernews_pb "github.com/easyCZ/grpc-web-hacker-news/server/proto"
	"google.golang.org/grpc/grpclog"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

func main() {
	router := chi.NewRouter()

	grpcServer := grpc.NewServer()
	hackernewsService := hackernews.NewHackerNewsService(nil)
	hackernews_pb.RegisterHackerNewsServiceServer(grpcServer, hackernewsService)

	wrappedGrpc := grpcweb.WrapServer(grpcServer)
	router.Use(func(handler http.Handler) http.Handler {
		return wrappedGrpc
	})

	router.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(":8900", router); err != nil {
		grpclog.Fatalf("failed starting http2 server: %v", err)
	}
}
