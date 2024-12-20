package helpers

import (
	"fmt"
	"net/http"

	l "github.com/peterszarvas94/goat/logger"
)

func Unauthorized(w http.ResponseWriter, r *http.Request, msg string, args ...any) {
	l.Logger.Error(msg, args...)
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintln(w, msg)
}
