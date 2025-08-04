package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Request Entered | PATH: %v | METHOD: %v\n", r.URL.Path, r.Method)

		next.ServeHTTP(w, r)

		fmt.Printf("Request Completed | PATH: %v | METHOD: %v | DURATION: %v\n", r.URL.Path, r.Method, time.Since(start))
	})
}
func SecurityHeaders(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Type", "application/json")
	c.Response().Header.Add("Content-Security-Policy", "default-src 'self'")
	c.Response().Header.Add("X-Frame-Options", "Deny")
	c.Response().Header.Add("X-XSS-Protection", "1; mode=block")
	return c.Next()
}
