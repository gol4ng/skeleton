package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Logging() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, handler)
	}
}
