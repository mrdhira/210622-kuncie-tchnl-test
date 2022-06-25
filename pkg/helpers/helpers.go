package helpers

import "github.com/gin-gonic/gin"

func ResponseHTTP(message string, error interface{}, data interface{}) map[string]any {
	return gin.H{
		"message": message,
		"error":   error,
		"data":    data,
	}
}
