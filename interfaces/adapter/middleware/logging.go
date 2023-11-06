package middleware

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string

	statusColor = param.StatusCodeColor()
	methodColor = param.MethodColor()
	resetColor = param.ResetColor()

	if param.Latency > time.Minute {
		param.Latency = param.Latency - param.Latency%time.Second
	}

	s := fmt.Sprintf("[GIN] %s %3d %s| %13v | %15s |%s %-7s %s %#v",
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
	)

	if param.ErrorMessage != "" {
		s += "\n" + param.ErrorMessage
	}

	return s
}

// LoggingMiddleware instance a Logger middleware.
func LoggingMiddleware() gin.HandlerFunc {
	conf := gin.LoggerConfig{}
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		log.Printf("%s", formatter(param))
	}
}
