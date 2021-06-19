package controller

import (
	"demo/common"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	log.Println("upload")
	err := ctx.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		panic(err)
	}
	log.Println("form")
	form, err := ctx.MultipartForm()
	if err != nil {
		FailureMsg(ctx, err.Error())
		return
	}
	for _, multifile := range form.File {
		for _, fd := range multifile {
			path := filepath.Join(common.UploadPath, fd.Filename)
			err := ctx.SaveUploadedFile(fd, path)
			if err != nil {
				log.Panicln("save err:", fd.Filename, err.Error())
				return
			}
			log.Println(fd.Filename)
		}
	}
	SuccessMsg(ctx, "chenggong")
}
