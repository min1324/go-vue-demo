package controller

import (
	"demo/common"
	"demo/model"
	"demo/validator"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	// get args
	var registerUser model.User
	err := ctx.Bind(&registerUser)
	fmt.Println(registerUser, err)
	// data valid
	if len(registerUser.Phone) != 11 {
		log.Println("phone:", registerUser.Phone)
		ResponeMsg(ctx, http.StatusUnprocessableEntity, 422, "phone must 11 bit.")
		return
	}
	if len(registerUser.Password) < 6 {
		ResponeMsg(ctx, http.StatusUnprocessableEntity, 422, "password at least 6 bit.")
		return
	}
	if len(registerUser.Name) == 0 {
		registerUser.Name = common.RandomString(10)
	}

	// check phone num
	if isExistsPhone(db, registerUser.Phone) {
		ResponeMsg(ctx, http.StatusUnprocessableEntity, 422, "user exist.")
		return
	}

	// create user
	hasPassword, err := bcrypt.GenerateFromPassword(
		[]byte(registerUser.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ResponeMsg(ctx, http.StatusInternalServerError, 500, "crypt err.")
		return
	}
	registerUser.Password = string(hasPassword)
	db.Create(&registerUser)
	token, err := common.CreateToken(&registerUser)
	if err != nil {
		ResponeMsg(ctx, http.StatusInternalServerError, 500, "system err.")
		return
	}

	// return
	Success(ctx, gin.H{"token": token}, "register success.")
}

func Login(ctx *gin.Context) {
	db := common.GetDB()
	// get args
	var loginUser model.User
	err := ctx.Bind(&loginUser)
	if err != nil {
		log.Println(err)
		ResponeMsg(ctx, http.StatusUnprocessableEntity, 422, "数据错误.")
		return
	}
	fmt.Println(loginUser.Phone, loginUser.Password)

	// data valider
	if len(loginUser.Phone) != 11 {
		log.Println("phone:", loginUser.Phone)
		ResponeMsg(ctx, http.StatusUnprocessableEntity, 422, "phone must 11 bit.")
		return
	}
	if len(loginUser.Password) < 6 {
		ResponeMsg(ctx, http.StatusUnprocessableEntity, 422, "password at least 6 bit.")
		return
	}

	// check phone
	var user model.User
	db.Where("phone=?", loginUser.Phone).First(&user)
	if user.ID == 0 {
		ResponeMsg(ctx, http.StatusUnprocessableEntity, 422, "user not exists.")
		return
	}
	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		FailureMsg(ctx, "password err.")
		return
	}

	// send token
	token, err := common.CreateToken(&user)
	if err != nil {
		ResponeMsg(ctx, http.StatusInternalServerError, 500, "system err.")
		log.Printf("token pares err:%s\n", err.Error())
		return
	}
	log.Println("login:", user.Phone)
	// return
	Success(ctx, gin.H{"token": token}, "login success.")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	log.Println("user:", validator.ToUserDto(user.(model.User)))
	Respone(ctx, http.StatusOK, 200, gin.H{"user": validator.ToUserDto(user.(model.User))}, "")
}

func isExistsPhone(db *gorm.DB, phone string) bool {
	var user model.User

	db.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
