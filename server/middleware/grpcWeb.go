package middleware

import (
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"net/http"
)

type GrpcWebMiddleware struct {
	*grpcweb.WrappedGrpcServer
}

func (m *GrpcWebMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsAcceptableGrpcCorsRequest(r) || m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewGrpcWebMiddleware(grpcWeb *grpcweb.WrappedGrpcServer) *GrpcWebMiddleware {
	return &GrpcWebMiddleware{grpcWeb}
}
