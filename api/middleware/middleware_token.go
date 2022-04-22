package middleware

import (
	"api/auth"
	"api/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

var UserId int

func MyMiddleware(service auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokens := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")[1]
		if tokens == "" {
			helper.Exception(ctx, false, "Unauthorized", nil)
			ctx.Abort()
			return
		}
		token, err := service.FindByToken(tokens)
		if token.UserId == 0 || token.CreatedAt.IsZero() || err != nil {
			helper.Exception(ctx, false, "Token is incorrect", nil)
			ctx.Abort()
			return
		}
		UserId = token.UserId

		if err != nil {
			helper.Exception(ctx, false, "Token is incorrect", nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
