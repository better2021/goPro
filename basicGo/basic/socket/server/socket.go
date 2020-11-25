package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server demo
func main()  {
	listen,err:=net.Listen("tcp","127.0.0.1:20000")
	if err!=nil{
		fmt.Println("listen to tcp failed",err)
		return
	}
	fmt.Println("服务端已连接")
	for {
		// 等待客户端来建立连接
		conn,err:=listen.Accept()
		if err!=nil{
			fmt.Println("accept failed",err)
			continue
		}
		// 启动一个单独的goroutine去处理连接
		go process(conn)

	}

}


func process(conn net.Conn){
	defer conn.Close() // 处理完通讯后关闭连接
	// 针对当前的连接做数据的发送和接受操作
	for {
		reader:=bufio.NewReader(conn)
		var buf [128]byte
		n,err:=reader.Read(buf[:])
		if err !=nil{
			fmt.Println("read from conn failed",err)
			break
		}
		recv:=string(buf[:n])
		fmt.Println("接收到的数据",recv)
		conn.Write(buf[:n])
	}
}