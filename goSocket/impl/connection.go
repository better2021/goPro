package impl

import "github.com/gorilla/websocket"

type Connection struct {
	wsConn *websocket.Conn
	inChan chan[]byte
	outChan chan[]byte
}

func InitConnection(wsConn *websocket.Conn)(conn *Connection,err error){
	conn = &Connection{
		wsConn:wsConn,
		inChan:make(chan []byte,1000),
		outChan:make(chan []byte,1000),
	}

	// 启动读协程
	go conn.readLoop()
	// 启动写携程
	go conn.writeLoop()
	return
}

func (conn *Connection) ReadMessage() (data []byte,err error) {
	data = <-conn.inChan
	return
}

func (conn *Connection) WriteMessage(data []byte) (err error) {
	conn.outChan <-data
	return
}

func (conn *Connection) Close() {
	conn.wsConn.Close()
}

// 内部实现
func (conn *Connection) readLoop(){
	var (
		data []byte
		err error
	)
	for{
		if _,data,err = conn.wsConn.ReadMessage();err!=nil{
			goto Err
		}
		conn.inChan <- data
	}
Err:
	conn.Close()
}

func (conn *Connection) writeLoop(){
	var (
		data []byte
		err error
	)
	for {
		data =<-conn.outChan
		if err = conn.wsConn.WriteMessage(websocket.TextMessage,data);err !=nil{
			goto Err
		}
	}
	Err:
		conn.Close()
}