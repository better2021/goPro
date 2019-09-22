package session

import (
	"sync"
	"time"
	"videoServer/api/dbops"
	"videoServer/api/defs"
	"videoServer/api/utils"
)

var sessionMap *sync.Map

func init(){
	sessionMap = &sync.Map{}
}

func nowInMilli() int64{
	return time.Now().UnixNano()/100000
}

func deleteExpiredSession(sid string){
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB(){
	r,err := dbops.RetrieveAllSessions()
	if err !=nil{
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key,ss)
		return true
	})
}

func GenerateNewSessionId(un string) string{
	id,_ := utils.NewUUID()
	ct := nowInMilli()
	ttl:= ct + 30 * 60 * 1000 // 过期时间

	ss := &defs.SimpleSession{Username:un,TTL:ttl}
	sessionMap.Store(id,ss)
	dbops.InserSession(id,ttl,un)
	return id
}

func IsSessionExpired(sid string) (string,bool) {
	ss ,ok := sessionMap.Load(sid)
	if ok{
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct{
			deleteExpiredSession(sid)
			return "",true
		}
		return ss.(*defs.SimpleSession).Username,false
	}
	return "",true
}