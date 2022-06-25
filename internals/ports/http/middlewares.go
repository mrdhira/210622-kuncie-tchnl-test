package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// commonHeaderValidation func
func commonHeaderValidation(ctx *gin.Context) {
	if ctx.Request.Header[http.CanonicalHeaderKey("X-User-ID")][0] == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "X-User-ID header is required"})
	}

	ctx.Next()
}
