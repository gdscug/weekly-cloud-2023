package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type ResponseObserver struct {
	http.ResponseWriter
	status      int
	headerWrote bool
}

func (l *ResponseObserver) Write(p []byte) (int, error) {
	if !l.headerWrote {
		l.WriteHeader(http.StatusOK)
	}
	n, err := l.ResponseWriter.Write(p)

	return n, err
}

func (l *ResponseObserver) WriteHeader(code int) {
	l.ResponseWriter.WriteHeader(code)
	if l.headerWrote {
		return
	}
	l.headerWrote = true
	l.status = code
}

func LogRequests(h http.Handler, infoLog *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o := &ResponseObserver{ResponseWriter: w}
		h.ServeHTTP(o, r)

		addr := r.RemoteAddr
		if i := strings.LastIndex(addr, ":"); i != -1 {
			addr = addr[:i]
		}
		infoLog.Printf("%s - - [%s] %q %d %q %q",
			addr,
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
			o.status,
			r.Referer(),
			r.UserAgent())
	})

}
