package middleware

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
)

type RecoveryHandlerLogger struct {
}

func (r *RecoveryHandlerLogger) Println(a ...interface{}) {
	fmt.Println(a...)
}

// Configure recovery middleware
func Recovery(debug bool) func(http.Handler) http.Handler {
	return handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(debug),
		//handlers.RecoveryLogger(&RecoveryHandlerLogger{}),
	)
}
