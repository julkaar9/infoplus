package server

import "github.com/gin-gonic/gin"

func SVGMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "image/svg+xml")
		c.Next()
	}
}
