package middlewares

import (
	"net/http"

	"github.com/fatih/color"
)

// LogRequest -> logs req info
func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		color.Yellow("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
