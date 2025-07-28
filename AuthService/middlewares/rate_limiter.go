package middlewares

import (
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

var limiter = rate.NewLimiter(rate.Every(1*time.Minute), 5)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many requests, please try again later.", http.StatusTooManyRequests)
			fmt.Println("Rate limit exceeded for request:", r.URL.Path)
			return
		}

		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}
