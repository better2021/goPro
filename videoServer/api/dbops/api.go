package dbops

import (
	"database/sql"
	"log"
	"time"
	"videoServer/api/defs"
	"videoServer/api/utils"
)

func AddUserCredential(loginName string,pwd string) error{
	stmtIns,err:= dbConn.Prepare("insert into users (login_name,pwd) values (?,?)")
	if err!=nil{
		return err
	}
	stmtIns.Exec(loginName,pwd)
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string,error) {
	stmtOut,err:= dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil{
		log.Printf("%s",err)
		return "",err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err!=nil&&err!= sql.ErrNoRows{
		return "",err
	}
	defer stmtOut.Close()

	return pwd,err
}

func DeleteUser(loginName string,pwd string) error {
	stmtDel,err:=dbConn.Prepare("delete from users where login_name=? and pwd=?")
	if err!=nil{
		log.Printf("DeleteUser error:%s",err)
		return err
	}

	_,err = stmtDel.Exec(loginName,pwd)
	if err!=nil{
		return nil
	}
	defer stmtDel.Close()
	return nil
}


func AddNewVidep(aid int,name string) (*defs.VideoInfo,error) {
	// create uuid
	vid,err := utils.NewUUID()
	if err != nil{
		return nil,err
	}

	t:= time.Now()
	ctime := t.Format("2006-01-02 15:04:05")
	stmtIns,err := dbConn.Prepare("insert into video_info (id,author_id,name,display_ctime) values (?,?,?,?)")
	if err!=nil{
		return nil,err
	}

	_,err = stmtIns.Exec(vid,aid,name,ctime)
	if err !=nil{
		return nil,err
	}

	res:= &defs.VideoInfo{Id:vid,AuthodId:aid,Name:name,DisplayCtime:ctime}
	defer stmtIns.Close()
	return res,nil
}

func AddNewComments(vid string,aid int,content string) error {
	id ,err := utils.NewUUID()
	if err!=nil{
		return err
	}
	stmtIns,err := dbConn.Prepare("insert into comments (id,video_id,author_id,content) values (?,?,?,?)")
	if err !=nil{
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err !=nil{
		return err
	}

	defer  stmtIns.Close()
	return nil
}

func ListComments(vid string,from, to int) ([]*defs.Comment,error){
	stmtOut,err := dbConn.Prepare(`select comments.id,users.Login_name,comments.content from comments
insert join users on comments.author_id = users.id where comments.video_id = ? and comments.time > FROM_UNIXTIME(?)
and comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows,err:= stmtOut.Query(vid,from,to)
	if err !=nil{
		return res,err
	}

	// 使用rows.Next()和rows.Scan() 获取到每一条数据
	for rows.Next(){
		var id,name,context string
		if err := rows.Scan(&id,&name,&context);err!=nil{
			return res,err
		}

		c := &defs.Comment{Id:id,VideoId:vid,Author:name,Content:context}
		res = append(res,c)
	}
	defer stmtOut.Close()

	return res,nil
}