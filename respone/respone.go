package respone

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
format
{
	code:200,
	data:xxx,
	msg:xxx,
}
*/

func Respone(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Respone(ctx, http.StatusOK, 200, data, msg)
}

func Fail(ctx *gin.Context, data gin.H, msg string) {
	Respone(ctx, http.StatusOK, 400, data, msg)
}
