package middleware

import (
	"demo/common"
	"demo/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const prefix = "Bearer "

// 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get authorization header
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, prefix) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[len(prefix):]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}
		// 验证通过后获取 userid.
		userId := claims.UserId
		db := common.GetDB()

		var user model.User
		db.First(&user, userId)

		// user
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			ctx.Abort()
			return
		}

		// 用户存在，将user信息写入上写文
		ctx.Set("user", user)
		ctx.Next()
	}
}
