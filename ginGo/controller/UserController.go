package controller

import (
	"fmt"
	"ginGo/common"
	"ginGo/model"
	"ginGo/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// 查询手机号
func isTelephoneExis(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 判断手机号和密码的长度是否正确
func isRight(telephone string,password string,ctx *gin.Context) bool{
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "手机号必须为11位",
		})
		return false
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 423,
			"msg":  "密码不能少于6位",
		})
		return false
	}
	return true
}

// 用户注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	isReturn := isRight(telephone,password,ctx)
	if !isReturn{
		return
	}
	// 如果名称没有传给一个随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, password, telephone)
	// 判断手机号是否存在
	if isTelephoneExis(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 424,
			"msg":  "用户已存在",
		})
		return
	}
	// 用户不存在就创建用户
	hasedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err !=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "加密发送错误",
		})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})

}

// 用户登录
func Login(ctx *gin.Context){
	DB := common.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	isReturn := isRight(telephone,password,ctx)
	if !isReturn{
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 424,
			"msg":  "用户不存在",
		})
		return
	}

	// 判断密码是否正确
	fmt.Println(user.Password,"--")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest,gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}
	// 发送token
	token,err := common.ReleaseToken(user)
	if err !=nil{
		ctx.JSON(http.StatusInternalServerError,gin.H{
			"code":500,
			"msg":"系统异常",
		})
		fmt.Println(err)
		return
	}
	// 返回结果
	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{"token":token},
		"msg":"登录成功",
	})
}

// 用户信息
func Info(ctx *gin.Context){
	user,_ := ctx.Get("user")

	ctx.JSON(http.StatusOK,gin.H{"code":200,"data":gin.H{"user":user}})
}