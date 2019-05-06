package middleware

import (
	"net/http"
	"time"

	"github.com/jakewright/muxinator"
)

// Logger is an interface that the standard log package implements
type Logger interface {
	Printf(format string, v ...interface{})
}

// Log returns a middleware that logs all requests
func Log(logger Logger) muxinator.Middleware {
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
