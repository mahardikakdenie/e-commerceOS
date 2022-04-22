package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Responses(ctx *gin.Context, status bool, message string, data interface{}, page int, page_size int) {
	// check value data have value or not
	if data == nil {
		data = gin.H{}
	}
	var pagination interface{}
	if page == 0 && page_size == 0 || page < 0 && page_size < 0 {
		pagination = gin.H{}
	} else {
		pagination = gin.H{
			"per_page":     page_size,
			"current_page": page,
		}
	}

	// response JSON
	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  status,
			"message": message,
		},
		"pagination": pagination,
		"data":       data,
	})
}

func Exception(ctx *gin.Context, status bool, message string, data interface{}) {
	// check value data have value or not
	if data == nil {
		data = gin.H{}
	}
	// response JSON
	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"status":  status,
			"message": message,
		},
	})
}
