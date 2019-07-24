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

func Recovery(debug bool) func(http.Handler) http.Handler {
	return handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(false),
		//handlers.RecoveryLogger(&RecoveryHandlerLogger{}),
		handlers.PrintRecoveryStack(debug),
	)
}
