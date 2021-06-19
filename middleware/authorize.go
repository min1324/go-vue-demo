package middleware

import (
	"demo/common"
	"demo/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const prefix = "Bearer "

// Authorization middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, prefix) {
			response(c, http.StatusUnauthorized, "权限不足")
			c.Abort()
			return
		}
		tokenString = tokenString[len(prefix):]

		token, cliams, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response(c, http.StatusUnauthorized, "权限不足")
			c.Abort()
			return
		}

		// get user id
		userId := cliams.UserId
		db := common.GetDB()

		var user model.User
		db.First(&user, userId)

		if user.ID == 0 {
			response(c, http.StatusUnauthorized, "权限不足")
			c.Abort()
			return
		}

		// set user info into head.
		c.Set("user", user)
		c.Next()
	}
}

func response(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"code": code, "msg": msg})
}
