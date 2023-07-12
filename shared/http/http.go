package http

import (
	"net/http"

	"github.com/campbel/terpcatalog/util/log"
)

func LogWrapper(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("request", "method", r.Method, "path", r.URL.Path, "remote", r.RemoteAddr, "email", r.Header.Get("X-Auth-Request-Email"))
		next.ServeHTTP(w, r)
	})
}

func FileHandler(contentType string, content []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write(content)
	})
}
