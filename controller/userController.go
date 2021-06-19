package controller

import (
	"demo/common"
	"demo/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()

	var registerUser model.User
	ctx.Bind(&registerUser)

	var user model.User
	db.Where("phone=?", registerUser.Phone).First(&user)
	if user.ID != 0 {
		// user exists
		ResponeMsg(ctx, http.StatusUnprocessableEntity, "user exists")
		return
	}
	if registerUser.Name == "" {
		registerUser.Name = common.RandomString(10)
	}

	// crypt passworld
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ResponeMsg(ctx, http.StatusInternalServerError, "crypt err")
		return
	}
	registerUser.Password = string(hasPassword)

	db.Create(&registerUser)

	// create token
	token, err := common.CreateToken(&registerUser)
	if err != nil {
		ResponeMsg(ctx, 500, "create token failed")
		log.Println(err.Error())
		return
	}
	Success(ctx, gin.H{"token": token}, "register success.")
}

func Login(ctx *gin.Context) {
	db := common.GetDB()

	var loginUser model.User
	err := ctx.Bind(&loginUser)
	if err != nil {
		ResponeMsg(ctx, http.StatusUnprocessableEntity, "数据错误")
		return
	}

	var user model.User
	db.Where("phone=?", loginUser.Phone).First(&user)
	if user.ID == 0 {
		ResponeMsg(ctx, http.StatusUnprocessableEntity, "user not exists")
		return
	}
	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		ResponeMsg(ctx, http.StatusInternalServerError, "password err.")
		return
	}

	// create token
	token, err := common.CreateToken(&user)
	if err != nil {
		ResponeMsg(ctx, http.StatusInternalServerError, "create token err")
		return
	}

	Success(ctx, gin.H{"token": token}, "login success.")
}
