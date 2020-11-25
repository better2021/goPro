package dbops

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init()  {
	dbConn,err = sql.Open("mysql","root:709463253@/video?charset=utf8&parseTime=True&loc=Loca")
	if err !=nil{
		panic(err.Error())
	}
}