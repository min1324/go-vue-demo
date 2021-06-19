package controller

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

func ResponeMsg(ctx *gin.Context, httpStatus int, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": httpStatus, "data": gin.H{}, "msg": msg})
}

func Success(ctx *gin.Context, data gin.H, msg string) {
	Respone(ctx, http.StatusOK, 200, data, msg)
}

func SuccessMsg(ctx *gin.Context, msg string) {
	Respone(ctx, http.StatusOK, 200, gin.H{}, msg)
}

func Failure(ctx *gin.Context, data gin.H, msg string) {
	Respone(ctx, http.StatusOK, 400, data, msg)
}

func FailureMsg(ctx *gin.Context, msg string) {
	Respone(ctx, http.StatusOK, 400, gin.H{}, msg)
}
