
package middleware

import (
	"log"
	"net/http"
	"time"
)


func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("started request to %s", r.URL)
		log.Printf("method: %s, body: %s", r.Method, r.Body)
		next.ServeHTTP(w, r)
		log.Printf("completed request to %s in %s", r.URL, time.Since(start))
	})
}