package controller

import (
	"ginGo/common"
	"ginGo/model"
	"ginGo/response"
	"ginGo/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost);err != nil{
		log.Println(err.Error())
		response.Fail(ctx,"分类名称必填",nil)
		return
	}

	// 获取登录用户 user
	user,_ := ctx.Get("user")

	// 创建文章
	post := model.Post{
		UserId:      user.(model.User).ID,
		Category_id: requestPost.Category_id,
		Title:       requestPost.Title,
		HeadImg:     requestPost.HeadImg,
		Content:     requestPost.Content,
	}

	if err := p.DB.Create(&post).Error;err!=nil{
		panic(err)
		return
	}
	response.Success(ctx,nil,"创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost);err != nil{
		log.Println(err.Error())
		response.Fail(ctx,"分类名称必填",nil)
		return
	}

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	if p.DB.Where("id=?",postId).First(&post).RecordNotFound(){
		response.Fail(ctx,"文章不存在",nil)
		return
	}

	// 获取用户登录 user
	user,_ := ctx.Get("user")
	userId := user.(model.User).ID

	// 判断当前用户是否为文章的作者
	if userId != post.UserId{
		response.Fail(ctx,"无权限，文章只有作者可以编辑",nil)
		return
	}

	// 更新文章
	if err := p.DB.Model(&post).Update(requestPost).Error;err != nil{
		response.Fail(ctx,"更新失败",nil)
		return
	}

	response.Success(ctx,gin.H{"post":post},"更新成功")
}

func (p PostController) Search(ctx *gin.Context) {
	// 获取path中的id
	var count int
	var post []model.Post
	if err:=p.DB.Preload("Category").First(&post).Count(&count).Error;err!=nil{
		response.Fail(ctx,"文章不存在",nil)
		return
	}

	response.Success(ctx,gin.H{"post":post,"count":count},"成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	// 判断文章是否存在
	if p.DB.Where("id = ?",postId).First(&post).RecordNotFound(){
		response.Fail(ctx,"文章不存在",nil)
		return
	}

	// 获取用户登录 user
	user,_ := ctx.Get("user")
	userId := user.(model.User).ID

	// 判断当前用户是否为文章的作者
	if userId != post.UserId{
		response.Fail(ctx,"无权限，文章只有作者可以编辑",nil)
		return
	}

	p.DB.Delete(&post)

	response.Success(ctx,gin.H{"post":post},"删除成功")
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB:db}
}

// 分页
func (p PostController) PageList(ctx *gin.Context){
	// 获取分页参数
	pageNum,_ := strconv.Atoi(ctx.DefaultQuery("pageNum","1"))
	pageSize,_ := strconv.Atoi(ctx.DefaultQuery("pageSize","20"))

	// 分页
	var posts []model.Post
	p.DB.Preload("Category").Order("created_at desc").Offset((pageNum - 1)*pageSize).Limit(pageSize).Find(&posts)

	// 记录的总条数
	var total int
	p.DB.Model(model.Post{}).Count(&total)

	response.Success(ctx,gin.H{"data":posts,"total":total},"success")
}
