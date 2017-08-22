package middleware

import (
    "log"
    "net/http"
    "time"
)

// This replicates the negroni.HandlerFunc type but decouples the code from the library
type Middleware func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)

func Log(logger *log.Logger) Middleware {
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
