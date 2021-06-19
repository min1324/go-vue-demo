package controller

import (
	"demo/model"
	"demo/repository"
	"demo/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryController interface {
	RestController
}

type CategoryContoller struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	// 自动迁移
	repository.DB.AutoMigrate(&model.Category{})
	return CategoryContoller{Repository: *repository}
}

func (c CategoryContoller) Create(ctx *gin.Context) {
	var requestCategory validator.CreateCategoryRequest
	err := ctx.ShouldBind(&requestCategory)
	if err != nil {
		FailureMsg(ctx, "数据验证错误，分类名称必填")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		FailureMsg(ctx, "创建失败")
		return
	}
	Success(ctx, gin.H{"category": category}, "创建成功")
}

func (c CategoryContoller) Update(ctx *gin.Context) {
	// 绑定 body 中的参数
	var requestCategory validator.CreateCategoryRequest
	err := ctx.ShouldBind(&requestCategory)
	if err != nil {
		FailureMsg(ctx, "数据验证错误，分类名称必填")
		return
	}
	// 获取 path 中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		FailureMsg(ctx, "分类不存在")
		return
	}
	// 更新分类
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		FailureMsg(ctx, "更新失败")
		return
	}
	Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryContoller) Show(ctx *gin.Context) {
	// 获取 path 中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		FailureMsg(ctx, "分类不存在")
		return
	}
	Success(ctx, gin.H{"category": category}, "success")
}

func (c CategoryContoller) Delete(ctx *gin.Context) {
	// 获取 path 中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.Repository.DeleteById(categoryId); err != nil {
		FailureMsg(ctx, "删除失败")
		return
	}
	// 直接删除
	SuccessMsg(ctx, "success")
}
