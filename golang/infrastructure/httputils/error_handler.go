package httputils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Logger receives gin contexts and return formated string with header and body
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		serviceTraceID := param.Request.Header.Get("Service-Traceid")
		return fmt.Sprintf("tracking-id: %s | method: %s | code: %d | ip: %s | endpoint: %s | latency: %s | agent: %s | date: %s \n",
			serviceTraceID,
			param.Method,
			param.StatusCode,
			param.ClientIP,
			param.Path,
			param.Latency,
			param.Request.UserAgent(),
			param.TimeStamp.Format("2006-01-02 15:04:05"),
		)
	})
}

//ErrorHandler func
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastError := c.Errors.ByType(gin.ErrorTypePrivate).Last()

		if lastError != nil {
			log.Println(lastError)
			c.JSON(http.StatusInternalServerError, gin.H{"error": lastError.Error()})
		}
	}
}
