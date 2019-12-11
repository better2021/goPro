package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Info struct {
	Id int64
	Name string
	Path string
	Note string
	CreateTime time.Time
}

// 定义数据库
var Db *sqlx.DB

func init(){
	db,err := sqlx.Open("mysql","root:709463253@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(db,err)
	if err != nil{
		log.Fatal(err)
	}
	err = db.Ping() // 查看数据库是否连接
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("数据库已连接")
	Db = db
}