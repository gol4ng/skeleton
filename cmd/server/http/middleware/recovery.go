package middleware

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type RecoveryHandlerLogger struct {
}

func (r *RecoveryHandlerLogger) Println(a ...interface{}) {
	fmt.Println(a...)
}

func Recovery(debug bool) mux.MiddlewareFunc {
	return handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(false),
		//handlers.RecoveryLogger(&RecoveryHandlerLogger{}),
		handlers.PrintRecoveryStack(debug),
	)
}
