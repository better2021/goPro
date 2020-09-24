package controller

import (
	"ginGo/common"
	"ginGo/model"
	"ginGo/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	if requestCategory.Name == ""{
		response.Fail(ctx,"分类名称必填",nil)
		return
	}

	c.DB.Create(&requestCategory)
	response.Success(ctx,gin.H{"data":requestCategory},"")
}

func (c CategoryController) Update(ctx *gin.Context) {
	panic("implement me")
}

func (c CategoryController) Search(ctx *gin.Context) {
	panic("implement me")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	panic("implement me")
}
