package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"basicGo/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
 数据库及连接配置
  return出db
*/

func Config() *gorm.DB {
	var db *gorm.DB
	var err error
	db,err = gorm.Open("mysql","root:709463253@/node?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println("数据库连接失败",err.Error())
	}else {
		fmt.Println("数据库已连接！")
		// 关联数据表自动迁移
		db.AutoMigrate(&model.User{})

		// 检查模型·User·是否存在
		hasTabUser := db.HasTable(&model.User{})
		fmt.Println(hasTabUser,"--")
		if !hasTabUser{
			// 如果没有User表，则为User模型创建User表，CHARSET=utf8设置数据库的字符集为utf8
			db.Set("gorm:table_options","ENGINE=InnoDB CHARSET=utf8").CreateTable(&model.User{})
		}
	}
	return db
}
