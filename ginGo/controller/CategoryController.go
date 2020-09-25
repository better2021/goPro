package controller

import (
	"fmt"
	"ginGo/common"
	"ginGo/model"
	"ginGo/response"
	"ginGo/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})

	return CategoryController{DB:db}
}

// 创建
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory);err != nil{
		response.Fail(ctx,"分类名称必填",nil)
		return
	}

	category := model.Category{Name:requestCategory.Name}
	c.DB.Create(&category)

	response.Success(ctx,gin.H{"data":requestCategory},"")
}

// 更新
func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中得参数
	var requestCategory model.Category
	err := ctx.Bind(&requestCategory)
	if err !=nil{
		fmt.Println(err.Error())
	}

	if requestCategory.Name == ""{
		response.Fail(ctx,"分类名称必填",nil)
		return
	}

	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id")) // 路由上传过来的id是字符串，需要转为int类型

	var updateCategory model.Category
	if c.DB.First(&updateCategory,categoryId).RecordNotFound(){
		response.Fail(ctx,"分类不存在",nil)
		return
	}

	// 更新分类 可以传三种类型 map ， struct ， name value
	c.DB.Model(&updateCategory).Update("name",requestCategory.Name)
	response.Success(ctx,gin.H{"data":updateCategory},"修改成功")

}

// 查找
func (c CategoryController) Search(ctx *gin.Context) {
	// 获取path中的参数
	var count int
	var category []model.Category
	if err:=c.DB.Find(&category).Count(&count).Error;err!=nil{
		response.Fail(ctx,"分类不存在",nil)
		return
	}

	response.Success(ctx,gin.H{"data":category,"count":count},"查询成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id")) // 路由上传过来的id是字符串，需要转为int类型

	if err := c.DB.Delete(model.Category{},categoryId).Error;err!=nil{
		response.Fail(ctx,"删除失败",nil)
		return
	}

	response.Success(ctx,nil,"删除成功")
}
