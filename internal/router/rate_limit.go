package router

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var visitors = make(map[string]time.Time)
var mu sync.Mutex

func RateLimit() gin.HandlerFunc {

	return func(c *gin.Context) {

		ip := c.ClientIP()

		mu.Lock()

		last, exists := visitors[ip]

		if exists && time.Since(last) < time.Second {
			mu.Unlock()

			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})

			c.Abort()
			return
		}

		visitors[ip] = time.Now()

		mu.Unlock()

		c.Next()
	}
}