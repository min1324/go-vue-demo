package controller

import (
	"demo/common"
	"demo/model"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()

	var registerUser model.User
	ctx.Bind(&registerUser)

	var user model.User
	db.Where("phone=?", registerUser.Phone).First(&user)
	if user.ID != 0 {
		// user exists
		FailureMsg(ctx, "user exists")
		return
	}
	if registerUser.Name == "" {
		registerUser.Name = common.RandomString(10)
	}
	db.Create(&registerUser)

	SuccessMsg(ctx, "register success.")
}

func Login(ctx *gin.Context) {

}
