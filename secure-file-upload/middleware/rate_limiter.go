package middleware

import (
	"net/http"
	"sync"
	"time"
)

var rateLimiter = make(map[string]int)
var mutex sync.Mutex

// RateLimitMiddleware enforces rate limiting per IP address
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		mutex.Lock()
		rateLimiter[ip]++
		mutex.Unlock()

		if rateLimiter[ip] > 10 {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		// Reset the rate limit after a minute
		go func() {
			time.Sleep(time.Minute)
			mutex.Lock()
			rateLimiter[ip] = 0
			mutex.Unlock()
		}()

		next.ServeHTTP(w, r)
	})
}
