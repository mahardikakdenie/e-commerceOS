package middleware

import (
	"api/auth"
	"api/customer_auth"
	"api/helper"

	"strings"

	"github.com/gin-gonic/gin"
)

var UserId uint
var CustomerId uint

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

func CustomerMiddleware(service customer_auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokens := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")[1]
		if tokens == "" {
			helper.Exception(ctx, false, "Unauthorized", nil)
			ctx.Abort()
			return
		}
		token, err := service.FindCustomerByToken(tokens)
		if err != nil {
			helper.Exception(ctx, false, "Token is incorrect", nil)
			ctx.Abort()
			return
		}
		if token.CustomerId == 0 || token.CreatedAt.IsZero() {
			helper.Exception(ctx, false, "Token is incorrect", nil)
			ctx.Abort()
			return
		}
		CustomerId = token.CustomerId
		ctx.Next()

	}
}
