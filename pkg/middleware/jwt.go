package middleware

import (
	"net/http"
	"strings"

	"github.com/Armageddon6026/zender/pkg/common"
	"github.com/Armageddon6026/zender/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthJwtRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.FullPath() == common.API_BASE_PATH+"/login" {
			ctx.Next()
			return
		}
		authorization := ctx.GetHeader("Authorization")
		jwtToken, ok := strings.CutPrefix(authorization, "Bearer ")
		if !ok {
			jwtToken, _ = ctx.Cookie(common.JWT_HEADER)
			if jwtToken == "" {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{
					Msg: "JWT token is not set",
				})
				return
			}
		}
		authJwtToken(jwtToken, ctx)
	}
}

func authJwtToken(token string, ctx *gin.Context) {
	tokenClaims, err := jwt.ParseWithClaims(token, &util.Claims{}, func(*jwt.Token) (any, error) {
		return util.GetJwtSecretByByte(), nil
	})
	if err != nil {
		message := err.Error()
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}

		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{Msg: message})
		return
	}

	if claims, ok := tokenClaims.Claims.(*util.Claims); !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{
			Msg: "JWT token payload is improper",
		})
		return
	} else {
		if claims.Account == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse{
				Msg: "JWT token payload is improper",
			})
			return
		}

		ctx.Set("account", claims.Account)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}
