package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main(){
	http.HandleFunc("/ws",WsHandler)
	http.ListenAndServe(":8080",nil)
}

func WsHandler(w http.ResponseWriter,r *http.Request){
	var (
		conn *websocket.Conn
		err error
		data []byte
	)
	if conn,err = upgrader.Upgrade(w,r,nil);err != nil{
		return
	}

	// websocket
	for {
		if _,data,err = conn.ReadMessage();err !=nil{
			goto Err
		}
		if err = conn.WriteMessage(websocket.TextMessage,data);err !=nil{
			goto Err
		}
	}
	Err:
		conn.Close()
}