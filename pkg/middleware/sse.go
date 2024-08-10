package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SSERequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")
		ctx.Writer.Header().Set("Transfer-Encoding", "chunked")

		if _, ok := ctx.Writer.(http.Flusher); !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Browser not support SSE"})
			return
		}

		ctx.Next()
	}
}
