package controller

import (
	"demo/common"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// ref: https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html
// @Summary Show an account
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path string true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} model.HTTPError
// @Router /accounts/{id} [get]
func Upload(c *gin.Context) {
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		panic(err)
	}
	form, err := c.MultipartForm()
	if err != nil {
		FailureMsg(c, err.Error())
		return
	}
	for _, multifile := range form.File {
		for _, fd := range multifile {
			path := filepath.Join(common.UploadPath, fd.Filename)
			err := c.SaveUploadedFile(fd, path)
			if err != nil {
				log.Panicln("save err:", fd.Filename, err.Error())
				return
			}
			log.Println(fd.Filename)
		}
	}
	SuccessMsg(c, "上传文件成功.")
}
