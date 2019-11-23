package models

import (
	"blogweb_gin/config"
	"blogweb_gin/database"
	"fmt"
)

type Article struct {
	Id int
	Title string
	Tags string
	Short string
	Content string
	Author string
	Createtime int64
	//Status int //Status=0为正常，1为删除，2为冻结
}

// 添加文章
func AddArticle(article Article)(int64,error){
	i,err := insertArticle(article)
	SetArticleRowsNum()
	return i,err
}

// 插入一篇文章
func insertArticle(article Article)(int64,error){
	return database.ModifyDB("insert into article(title,tags,short,content,author,createtime) values(?,?,?,?,?,?)",
		article.Title,article.Tags,article.Short,article.Author,article.Createtime)
}

/*
* 查询文章
*/
// 根据页码查询文章
func FindArticleWithPage(page int) ([]Article,error){
	page --
	fmt.Println("-->page",page)
	// 从配置文件中获取每页的文章数量
	return QueryArticleWithPage(page,config.NUM)
}

/**
分页查询数据库
limit分页查询语句，
	语法：limit m，n

	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据

注意limit前面咩有where
*/
func QueryArticleWithPage(page,num int) ([]Article,error){
	sql := fmt.Sprintf("limit %d,%d",page*num,num)
	return QueryUserWightCon(sql)
}

func QueryArticlesWithCon(sql string) ([]Article,error) {
	Sql := "select id,title,tags,short,content,author,createtime from article" +sql
	rows,err := database.QueryDB(Sql)
	if err != nil{
		return nil,err
	}
	var artList []Article
	for rows.Next(){
		id := 0
		title := ""
		short := ""
		tags := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id,&title,&tags,&short,&content,&author,&createtime)
		art := Article{id,title,tags,short,content,author,createtime}
		artList = append(artList,art)
	}
	return artList,nil
}

/*
* 翻页
*/
//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0
func GetArtcileRowsNum() int{
	if artcileRowsNum == 0{
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

