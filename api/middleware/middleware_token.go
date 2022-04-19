package middleware

import (
	"api/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var UserId int

func MyMiddleware(service auth.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokens := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")[1]
		if tokens == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"meta": gin.H{
					"status":  false,
					"message": "Unauthorized",
				},
				"data": gin.H{},
			})
			ctx.Abort()
			return
		}
		token, err := service.FindByToken(tokens)
		if token.UserId == 0 || token.CreatedAt.IsZero() || err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"meta": gin.H{
					"status":  false,
					"message": "Token is incorrect",
				},
				"data": gin.H{},
			})
			ctx.Abort()
			return
		}
		UserId = token.UserId
		// ctx.Next()

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
	}
}
