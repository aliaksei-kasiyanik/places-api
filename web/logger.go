package web

import (
	"log"
	"os"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

type Logger struct {
	negroni.ALogger
}

func NewLogger() *Logger {
	return &Logger{log.New(os.Stdout, "[place-api] ", 0)}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	l.Printf(" %s\t%s\t%v\t%s\t%v",
		r.Method,
		r.RequestURI,
		res.Status(),
		http.StatusText(res.Status()),
		time.Since(start))
}
