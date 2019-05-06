package log

import (
	"net/http"

	"github.com/jakewright/drawbridge/log"
	"github.com/jakewright/drawbridge/plugin"
	"github.com/jakewright/muxinator"
)

func init() {
	plugin.RegisterPlugin("log", &Logger{})
}

type Logger struct{}

func (l *Logger) Middleware(map[string]interface{}) (muxinator.Middleware, error) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next(w, r)
	}, nil
}
