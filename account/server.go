package account

import (
	"context"
	"net/http"
)

func NewHTTPServer (ctx context.Context, endpoint Endpoint) http.Handler {

}

func commonMiddleware (next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
