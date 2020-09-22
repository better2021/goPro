package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
  	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct{
	gorm.Model
	Name string	`gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

func main(){
	db := InitDB()
	defer db.Close()

	r := gin.Default()
	r.POST("/api/auth/register",func(ctx *gin.Context){
		// 获取参数
		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")
		// 数据验证
		if len(telephone)!=11{
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":422,
				"msg":"手机号必须为11位",
			})
			return
		}

		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":423,
				"msg":"密码不能少于6位",
			})
			return
		}
		// 如果名称没有传给一个随机字符串
		if len(name) == 0{
			name = RandomString(10)
		}
		log.Println(name,password,telephone)
		// 判断手机号是否存在
		if isTelephoneExis(db,telephone){
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":424,
				"msg":"用户已存在",
			})
			return
		}
		// 用户不存在就创建用户
		newUser := User{
			Name:name,
			Telephone:telephone,
			Password:password,
		}
		db.Create(&newUser)

		// 返回结果
		ctx.JSON(http.StatusOK,gin.H{
			"msg":"注册成功",
		})
		
	})
	r.Run()
}

// 导出随机字符串
func RandomString(n int) string{
	var letters = []byte("asdfghjklzxcvbnmASDFGHJKLZCVBNM")
	result := make([]byte,n)

	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 查询手机号
func isTelephoneExis(db *gorm.DB,telephone string) bool{
	var user User
	db.Where("telephone=?",telephone).First(&user)
	if user.ID != 0{
		return true
	}
	return false
}

// 初始化数据库
func InitDB() *gorm.DB{
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginGO"
	username := "root"
	password := "709463253"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
	username,
	password,
	host,
	port,
	database,
	charset)

	db,err := gorm.Open(driverName,args)
	if err != nil{
		panic("failed to connect database,err" + err.Error())
	}
	db.AutoMigrate(&User{})

	return db
}