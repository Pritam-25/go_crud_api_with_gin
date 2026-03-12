package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger logs each request with a short unique ID, method, path, status, and latency.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := fmt.Sprintf("%x", start.UnixNano())

		c.Set("requestID", requestID)
		c.Header("X-Request-ID", requestID)

		c.Next()

		log.Printf("[%s] %s %s %d %s",
			requestID[:8],
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			time.Since(start),
		)
	}
}
