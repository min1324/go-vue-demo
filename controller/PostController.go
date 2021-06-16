package controller

import (
	"demo/common"
	"demo/model"
	"demo/respone"
	"demo/validator"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return &PostController{DB: db}

}

func (p *PostController) Create(ctx *gin.Context) {
	var requestPost validator.CreatePostRequest
	// 数据验证
	err := ctx.ShouldBind(&requestPost)
	if err != nil {
		respone.Fail(ctx, nil, "数据验证错误")
		panic(err)
	}
	// 获取登陆用户
	user, _ := ctx.Get("user")

	// 创建文章
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Tittle:     requestPost.Tittle,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	// 插入数据

	if err := p.DB.Create(&post).Error; err != nil {
		respone.Fail(ctx, nil, "创建失败")
		return
	}
	respone.Success(ctx, nil, "创建成功")
}

func (p *PostController) Update(ctx *gin.Context) {
	var requestPost model.Post
	err := ctx.ShouldBind(&requestPost)
	if err != nil {
		respone.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	// 获取 path 中的 id
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Where("id=?", postId).First(&post).Error == gorm.ErrRecordNotFound {
		respone.Fail(ctx, nil, "文章不存在.")
		return
	}
	// 判断当前用户是否为文章的作者

	// 获取登陆用户
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		respone.Fail(ctx, nil, "文字不是你的，请勿修改")
		return
	}
	log.Println(post, requestPost)
	tx := p.DB.Model(&post).Updates(requestPost)
	err = tx.Error
	if err != nil {
		respone.Fail(ctx, nil, "更新失败")
		return
	}
	respone.Success(ctx, gin.H{"post": post}, "更新成功")

}

func (p *PostController) Show(ctx *gin.Context) {
	// 获取 path 中的 id
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Preload("Category").Where("id=?", postId).First(&post).Error == gorm.ErrRecordNotFound {
		respone.Fail(ctx, nil, "文章不存在.")
		return
	}
	respone.Success(ctx, gin.H{"post": post}, "成功")
}

func (p *PostController) Delete(ctx *gin.Context) {
	var requestPost validator.CreatePostRequest
	err := ctx.ShouldBind(&requestPost)
	if err != nil {
		respone.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return
	}

	// 获取 path 中的 id
	postId := ctx.Params.ByName("id")
	var post model.Post
	if p.DB.Where("id=?", postId).First(&post).Error == gorm.ErrRecordNotFound {
		respone.Fail(ctx, nil, "文章不存在.")
		return
	}
	// 判断当前用户是否为文章的作者

	// 获取登陆用户
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		respone.Fail(ctx, nil, "文字不是你的，请勿修改")
		return
	}
	if err := p.DB.Model(&post).Delete(requestPost).Error; err != nil {
		respone.Fail(ctx, nil, "删除失败")
		return
	}
	respone.Success(ctx, nil, "成功")
}

func (p PostController) PageList(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	PageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页

	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * PageSize).Limit(PageSize).Find(&posts)

	// 需要知道总条数
	var total int64
	p.DB.Model(&model.Post{}).Preload("Category").Count(&total)

	respone.Success(ctx, gin.H{"data": posts, "total": total}, "success")
}
