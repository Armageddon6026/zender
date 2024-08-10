package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Msg string `json:"msg,omitempty"`
}

func ApiNotFound(ctx *gin.Context) {
	if strings.HasPrefix(ctx.Request.URL.Path, "/api") {
		ctx.JSON(http.StatusNotFound, errorResponse{Msg: "api route NotFound"})
		ctx.Abort()
		return
	}
	ctx.Next()
}
