package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func GrpcRequest(g *grpc.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Is HTTP/2 ?
		// Is grpc request ?
		if ctx.Request.ProtoMajor == 2 && strings.HasPrefix(ctx.GetHeader("Content-Type"), "application/grpc") {
			ctx.Status(http.StatusOK)
			g.ServeHTTP(ctx.Writer, ctx.Request)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
