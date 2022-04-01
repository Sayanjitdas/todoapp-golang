package middleware

import (
	"log"
	"net/http"
)

type Logger struct {
	handler http.Handler
}

//ServeHTTP is the handler which pass
//the request to actual handler wrapped with logging statement
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

//logger middleware
func LoggerMiddleware(Requesthandler http.Handler) *Logger {
	logger := Logger{
		handler: Requesthandler,
	}
	return &logger
}
