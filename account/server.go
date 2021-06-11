package account

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPServer (ctx context.Context, endpoint Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
			endpoint.CreateUser,
			decodeUserRequest,
			encodeResponse,
		))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
			endpoint.GetUser,
			decodeEmailRequest,
			encodeResponse,
		))

	return r
}

func commonMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
