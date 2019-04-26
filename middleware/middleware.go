package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/jakewright/muxinator"
)

// Log returns a middleware that logs all requests
func Log(logger *log.Logger) muxinator.Middleware {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		start := time.Now()
		next(w, r)

		logger.Printf(
			"%s %s %s %s",
			start.Format("2016/01/02 15:04:05"),
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}
