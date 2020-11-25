package controllers

import (
	"blogweb_gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
当访问/add路径的时候回触发AddArticleGet方法
响应的页面是通过HTML
*/
func AddArticleGet(c *gin.Context){

	// 获取session
	islogin := GetSession(c)
	c.HTML(http.StatusOK,"write_article.html",gin.H{
		"IsLogin":islogin,
	})
}

func AddArticlePost(c *gin.Context) {
	//获取浏览器传输的数据，通过表单的name属性获取值
	//获取表单信息
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	author := c.PostForm("author")
	fmt.Sprintf("title:%s,tags:%s\n",title,tags)

	// 实例化model，将它写入数据库
	art := models.Article{0,title,tags,short,content,author,time.Now().Unix()}
	_,err := models.AddArticle(art)

	// 返回数据给浏览器
	response := gin.H{}
	if err == nil{
		// 成功
		response = gin.H{"code":1,"message":"ok"}
	}else {
		response = gin.H{
			"code":0,
			"message":"error",
		}
	}
	c.JSON(http.StatusOK,response)
}