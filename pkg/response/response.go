package response

import (
	"github.com/NhatHaoDev3324/GoTemplate/constant"
	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, status constant.SuccessStatus, message string, data interface{}) {
	ctx.JSON(int(status), gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

func Fail(ctx *gin.Context, status constant.FailStatus, message string) {
	ctx.JSON(int(status), gin.H{
		"success": false,
		"message": message,
	})
}
