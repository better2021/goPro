package main

import (
	"ginGo/common"
	"ginGo/route"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()	// 初始化数据库
	db := common.GetDB()
	defer db.Close()

	r := gin.Default()
	r = route.CollectRouter(r)
	r.Run()
}
