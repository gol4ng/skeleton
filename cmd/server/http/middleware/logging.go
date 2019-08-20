package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// Configure logging middleware
func Logging() func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return handlers.LoggingHandler(os.Stdout, handler)
	}
}
