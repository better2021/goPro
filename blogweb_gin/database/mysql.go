package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"  // mysql驱动（必须有）
)

var db *sql.DB

func InitMysql()  {
	fmt.Println("InitMysql...")
	if db == nil{
		db,_ = sql.Open("mysql","root:709463253@tcp(127.0.0.1:3306)/blogweb_gin?charset=utf8&parseTime=True&loc=Local")
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}
}

// 操作数据库
func ModifyDB(sql string,arg ...interface{})(int64,error){
	result,err := db.Exec(sql,arg...)
	if err != nil{
		log.Println(err)
		return 0,err
	}
	count,err := result.RowsAffected()
	if err != nil{
		log.Println(err)
		return 0,err
	}

	return count,nil
}

// 创建用户表
func CreateTableWithUser(){
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
	)`

	ModifyDB(sql)
}

// 查询(执行一个预期最多只会返回一个数据行的查询)
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

//  创建文章表
func CreateTableWithArticle(){
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
	)`
	ModifyDB(sql)
}

// 查询（执行一个会返回数据行的查询， 通常是一个 SELECT）
func QueryDB(sql string) (*sql.Rows,error){
	return db.Query(sql)
}

// ----图片----
func CreateTableWithAlbum()  {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
	)`
	ModifyDB(sql)
}



