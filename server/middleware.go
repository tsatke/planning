package server

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		bodySize := c.Writer.Size()

		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		event := s.log.Debug()
		if statusCode >= 500 && statusCode < 600 {
			event = s.log.Error()
		}
		event.
			Int("status", statusCode).
			Str("dur", duration.String()).
			Str("from", clientIP).
			Str("method", method).
			Int("respsz", bodySize).
			Str("path", path).
			Msg("serve request")
	}
}
